package queries

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	FindAll struct {
	}
	FindAllHandler struct {
		notification domain.NotificationRepository
	}
)

func NewFindAllHandler(notification domain.NotificationRepository) FindAllHandler {
	return FindAllHandler{notification: notification}
}

func (h FindAllHandler) FindAll(ctx context.Context) ([]domain.Notification, error) {
	return h.notification.FindAll(ctx)
}
