package commands

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	UpdateNotificationCommand struct {
		Id             string
		Subject        string
		Content        string
		SenderName     string
		SenderEmail    string
		RecipientEmail string
		RecipientName  string
		Status         domain.EmailStatus
		Type           domain.NotifiationType
	}

	UpdateNotificationHandler struct {
		notification domain.NotificationRepository
	}
)

func NewUpdateNotificationHandler(notification domain.NotificationRepository) UpdateNotificationHandler {
	return UpdateNotificationHandler{notification: notification}
}

func (h UpdateNotificationHandler) UpdateNotification(ctx context.Context, cmd UpdateNotificationCommand) (*domain.Notification, error) {
	notification := domain.Notification{
		Subject:        cmd.Subject,
		Content:        cmd.Content,
		SenderName:     cmd.SenderName,
		SenderEmail:    cmd.SenderEmail,
		RecipientEmail: cmd.RecipientEmail,
		RecipientName:  cmd.RecipientName,
		Status:         cmd.Status,
		Type:           cmd.Type,
	}

	err := h.notification.Update(ctx, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
