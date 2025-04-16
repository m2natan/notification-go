package queries

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	FindByStatus struct {
		Status domain.EmailStatus
	}

	FindByStatusHandler struct {
		notification domain.NotificationRepository
	}
)

func NewFindByStatusHandler(notification domain.NotificationRepository) FindByStatusHandler {
	return FindByStatusHandler{notification: notification}
}

func (h FindByStatusHandler) FindByStatus(ctx context.Context, cmd FindByStatus) ([]domain.Notification, error) {
	return h.notification.FindByStatus(ctx, cmd.Status)
}
