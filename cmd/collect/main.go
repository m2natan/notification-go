package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/IBM/sarama"
	"github.com/Kifiya-Financial-Technology/Notification-Service/handler"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/logging"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/postgres"
	"github.com/Kifiya-Financial-Technology/Notification-Service/notificationpb"
	"github.com/Kifiya-Financial-Technology/Notification-Service/nswagger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pg "gorm.io/driver/postgres"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

const (
	topic   = "notifications"
	groupID = "notification-service"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ No .env file found or failed to load it")
	} else {
		log.Println("✅ .env file loaded")
	}

	port := os.Getenv("REST_PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")

	grpcAddress := fmt.Sprintf("0.0.0.0:%s", grpcPort)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect to DB: %v", err)
	}
	fmt.Println("✅ Database connected")

	// Enable uuid-ossp extension
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("❌ failed to enable uuid-ossp extension: %v", err)
	}

	// Auto-migrate the Product model
	if err := db.AutoMigrate(&domain.Notification{}); err != nil {
		log.Fatalf("❌ failed to auto-migrate: %v", err)
	}
	fmt.Println("✅ Auto-migration complete")

	// Setup gRPC server
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("❌ failed to listen on port %v: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()

	// Dependency injection
	repo := postgres.NewNotificationRepository(db)
	app := application.New(repo)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app_logger := logging.LogApplicationAccess(app, logger)

	// Register gRPC handlers
	handler.NewServer(app_logger, grpcServer)

	fmt.Println("🚀 gRPC server running on :" + grpcAddress)

	// Setup gRPC-Gateway and Swagger
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err = notificationpb.RegisterNotificationServiceHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("❌ failed to register gRPC handlers: %v", err)
	}

	// Combine gRPC Gateway and Swagger handler into one HTTP server
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	httpMux.Handle("/swagger/", nswagger.SwaggerHandler())

	fmt.Println("🌐 HTTP server (REST + Swagger) running on :" + port)
	go func() {
		if err := http.ListenAndServe(":"+port, httpMux); err != nil {
			log.Fatalf("❌ HTTP server failed: %v", err)
		}
	}()

	// Initialize Kafka reader
	kafkaReader := []string{"kafka1:9091"}
	consumer, err := sarama.NewConsumer(kafkaReader, nil)
	if err != nil {
		log.Fatalf("❌ failed to connect to Kafka broker: %v", err)
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("❌ failed to get list of partitions: %v", err)
	}

	messages := make(chan *sarama.ConsumerMessage, 256)
	initialOffset := sarama.OffsetOldest //offset to start reading message from
	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			log.Printf("❌ Failed to consume partition %d: %v", partition, err)
			continue // Don't continue blindly if failed
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.AsyncClose()

			for message := range pc.Messages() {
				var notification domain.Notification

				err := json.Unmarshal(message.Value, &notification)
				if err != nil {
					log.Printf("❌ Failed to unmarshal Kafka message: %v", err)
					fmt.Printf("⚡ Raw message: %s\n", string(message.Value))
					continue // Don't break — just skip this bad message!
				}
				// Successfully parsed message
				app_logger.CreateNotification(context.Background(), commands.CreateNotificationCommand{
					Subject:       notification.Subject,
					Content:       notification.Content,
					SenderName:    notification.SenderName,
					Sender:        notification.Sender,
					Recipient:     notification.Recipient,
					RecipientName: notification.RecipientName,
					Type:          notification.Type,
				})
				messages <- message
			}
		}(pc)
	}
	// Run the gRPC server
	log.Fatal(grpcServer.Serve(listener))
}
