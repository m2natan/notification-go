package commands

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	DeleteNotificationCommand struct {
		Id string
	}

	DeleteNotificationHandler struct {
		notification domain.NotificationRepository
	}
)

func NewDeleteNotificationHandler(notification domain.NotificationRepository) DeleteNotificationHandler {
	return DeleteNotificationHandler{notification: notification}
}

func (h DeleteNotificationHandler) DeleteNotification(ctx context.Context, cmd DeleteNotificationCommand) error {
	return h.notification.Delete(ctx, cmd.Id)
}
