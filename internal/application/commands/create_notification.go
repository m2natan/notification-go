package commands

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	CreateNotificationCommand struct {
		Subject        string
		Content        string
		SenderName     string
		Sender    string
		Recipient string
		RecipientName  string
		Status         domain.EmailStatus
		Type           domain.NotificationType
	}

	CreateNotificationHandler struct {
		notification domain.NotificationRepository
	}
)

func NewCreateNotificationHandler(notification domain.NotificationRepository) CreateNotificationHandler {
	return CreateNotificationHandler{notification: notification}
}

func (h CreateNotificationHandler) CreateNotification(ctx context.Context, cmd CreateNotificationCommand) (*domain.Notification, error) {
	notification, err := domain.CreateNotification(cmd.Subject, cmd.Content, cmd.SenderName, cmd.Sender, cmd.Recipient, cmd.RecipientName, cmd.Status, cmd.Type)
	if err != nil {
		return nil, err
	}

	err = h.notification.Create(ctx, notification)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
