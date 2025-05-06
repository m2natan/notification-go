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
	"time"

	"github.com/IBM/sarama"
	"github.com/Kifiya-Financial-Technology/Notification-Service/handler"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/logging"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/notifier/push"
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
	topic         = "notifications"
	consumerGroup = "notification-service-group" // Use consumer groups for scalability
	numPartitions = 3                            // Match your Kafka topic partitions
	batchSize     = 100                          // Process logs in batches
	flushInterval = 1 * time.Second              // Max delay before flushing a batch
	maxRetries    = 3                            // Retry failed messages
	kafkaBrokers  = "kafka1:9091"
)

func init_database(dbHost, dbPort, dbUser, dbPass, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect to DB: %v", err)
		return nil, err
	}
	fmt.Println("✅ Database connected")

	// Enable uuid-ossp extension
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("❌ failed to enable uuid-ossp extension: %v", err)
		return nil, err
	}

	// Auto-migrate the Product model
	if err := db.AutoMigrate(&domain.Notification{}); err != nil {
		log.Fatalf("❌ failed to auto-migrate: %v", err)
		return nil, err
	}
	fmt.Println("✅ Auto-migration complete")
	return db, nil
}

func init_grpc_client(grpcPort string, grpcAddress string) (net.Listener, *grpc.Server, error) {

	// Setup gRPC server
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("❌ failed to listen on port %v: %v", grpcPort, err)
		return nil, nil, err
	}

	grpcServer := grpc.NewServer()
	return listener, grpcServer, nil
}

func main() {

	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ No .env file found or failed to load it")
	} else {
		log.Println("✅ .env file loaded")
	}

	// Database connection setup
	port := os.Getenv("REST_PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")

	db, err := init_database(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatalf("❌ failed to connect to DB: %v", err)
	}
	grpcAddress := fmt.Sprintf("0.0.0.0:%s", grpcPort)

	listener, grpcServer, err := init_grpc_client(grpcPort, grpcAddress)
	if err != nil {
		log.Fatalf("❌ failed to listen on port %v: %v", grpcPort, err)
	}

	// Dependency injection
	repo := postgres.NewNotificationRepository(db)
	app := application.New(repo)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app_logger := logging.LogApplicationAccess(app, logger)

	// Register gRPC handlers
	handler.NewServer(app_logger, grpcServer)

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
	httpMux.HandleFunc("/event", push.SSEHandler)

	// Start the HTTP server for REST + Swagger
	fmt.Println("🌐 HTTP server (REST + Swagger) running on :" + port)
	go func() {
		if err := http.ListenAndServe(":"+port, httpMux); err != nil {
			log.Fatalf("❌ HTTP server failed: %v", err)
		}
	}()

	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Fetch.Min = 1 << 20     // 1MB
	config.Consumer.Fetch.Default = 5 << 20 // 5MB
	config.ChannelBufferSize = 1024         // Buffer size

	// Initialize Kafka reader
	consumer, err := sarama.NewConsumer([]string{kafkaBrokers}, config)
	if err != nil {
		log.Fatalf("❌ failed to connect to Kafka broker: %v", err)
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("❌ failed to get list of partitions: %v", err)
	}

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Printf("❌ Failed to consume partition %d: %v", partition, err)
			continue // Don't continue blindly if failed
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.AsyncClose()

			for message := range pc.Messages() {
				var notification domain.Notification

				if err := json.Unmarshal(message.Value, &notification); err != nil {
					// If unmarshalling fails, print error and continue to the next message
					fmt.Printf("⚡ Invalid Kafka message: %s\n", string(message.Value))
					continue
				}
				// Successfully parsed message
				_, err := app_logger.CreateNotification(context.Background(), commands.CreateNotificationCommand{
					Subject:       notification.Subject,
					Content:       notification.Content,
					SenderName:    notification.SenderName,
					Sender:        notification.Sender,
					Recipient:     notification.Recipient,
					RecipientName: notification.RecipientName,
					Type:          notification.Type,
				})
				if err != nil {
					log.Printf("⚠️ Failed to process log: %v\n", err)
				} else {
					// Successfully processed log
					log.Printf("✅ Log processed: %s\n", notification) // print the actual message, not the raw value
				}
			}
		}(pc)
	}

	// Run the gRPC server
	log.Fatal(grpcServer.Serve(listener))
}
