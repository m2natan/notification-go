package handler

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/commands"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/application/queries"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/converter"
	"github.com/Kifiya-Financial-Technology/Notification-Service/notificationpb"
	"google.golang.org/grpc"
)

type NotificationServiceHandler struct {
	app application.App
	notificationpb.UnimplementedNotificationServiceServer
}

func NewServer(app application.App, grpcServer *grpc.Server) {
	notification := &NotificationServiceHandler{app: app}
	notificationpb.RegisterNotificationServiceServer(grpcServer, notification)
}

func (n *NotificationServiceHandler) CreateNotification(ctx context.Context, req *notificationpb.CreateNotificationRequest) (*notificationpb.Notification, error) {
	cmd := commands.CreateNotificationCommand{
		Subject:        req.Subject,
		Content:        req.Content,
		SenderName:     req.SenderName,
		Sender:    req.Sender,
		Recipient: req.Recipient,
		RecipientName:  req.RecipientName,
		Type:           converter.ConvertPbTypeToDomain(req.Type),
	}

	notification, err := n.app.CreateNotification(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return converter.ConvertNotificationToPb(notification), nil
}

func (n *NotificationServiceHandler) UpdateNotification(ctx context.Context, req *notificationpb.UpdateNotificationRequest) (*notificationpb.Notification, error) {
	cmd := commands.UpdateNotificationCommand{
		Id:             req.Id,
		Subject:        req.Subject,
		Content:        req.Content,
		SenderName:     req.SenderName,
		Sender:    req.Sender,
		Recipient: req.Recipient,
		RecipientName:  req.RecipientName,
		Type:           converter.ConvertPbTypeToDomain(req.Type),
	}

	notification, err := n.app.UpdateNotification(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return converter.ConvertNotificationToPb(notification), nil
}

func (n *NotificationServiceHandler) DeleteNotification(ctx context.Context, req *notificationpb.DeleteNotificationRequest) (*notificationpb.Notification, error) {
	cmd := commands.DeleteNotificationCommand{
		Id: req.Id,
	}
	query := queries.FindById{Id: req.Id}
	notification, err := n.app.FindById(ctx, query)
	if err != nil {
		return nil, err
	}
	err = n.app.DeleteNotification(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return converter.ConvertNotificationToPb(notification), nil
}

func (n *NotificationServiceHandler) GetNotificationById(ctx context.Context, req *notificationpb.GetNotificationByIdRequest) (*notificationpb.Notification, error) {
	query := queries.FindById{Id: req.Id}
	notification, err := n.app.FindById(ctx, query)
	if err != nil {
		return nil, err
	}
	return converter.ConvertNotificationToPb(notification), nil
}

func (n *NotificationServiceHandler) GetNotificationsByType(ctx context.Context, req *notificationpb.GetNotificationsByTypeRequest) (*notificationpb.GetNotificationsByTypeResponse, error) {
	query := queries.FindByType{Type: converter.ConvertPbTypeToDomain(req.Type)}
	notification, err := n.app.FindByType(ctx, query)
	if err != nil {
		return nil, err
	}
	var pbNotifications []*notificationpb.Notification
	for _, notif := range notification {
		pbNotifications = append(pbNotifications, converter.ConvertNotificationToPb(&notif))
	}

	return &notificationpb.GetNotificationsByTypeResponse{
		Notifications: pbNotifications,
	}, nil
}

func (n *NotificationServiceHandler) GetNotificationsByStatus(ctx context.Context, req *notificationpb.GetNotificationsByStatusRequest) (*notificationpb.GetNotificationsByStatusResponse, error) {
	query := queries.FindByStatus{Status: converter.ConvertPbStatusToDomain(req.Status)}
	notification, err := n.app.FindByStatus(ctx, query)
	if err != nil {
		return nil, err
	}
	var pbNotifications []*notificationpb.Notification
	for _, notif := range notification {
		pbNotifications = append(pbNotifications, converter.ConvertNotificationToPb(&notif))
	}

	return &notificationpb.GetNotificationsByStatusResponse{
		Notifications: pbNotifications,
	}, nil
}

func (n *NotificationServiceHandler) GetNotifications(ctx context.Context, req *notificationpb.GetNotificationsRequest) (*notificationpb.GetNotificationsResponse, error) {
	notification, err := n.app.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var pbNotifications []*notificationpb.Notification
	for _, notif := range notification {
		pbNotifications = append(pbNotifications, converter.ConvertNotificationToPb(&notif))
	}

	return &notificationpb.GetNotificationsResponse{
		Notifications: pbNotifications,
	}, nil
}
