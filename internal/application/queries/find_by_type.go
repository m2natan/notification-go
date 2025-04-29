package queries

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	FindByType struct {
		Type domain.NotificationType
	}

	FindByTypeHandler struct {
		notification domain.NotificationRepository
	}
)

func NewFindByTypeHandler(notification domain.NotificationRepository) FindByTypeHandler {
	return FindByTypeHandler{notification: notification}
}

func (h FindByIdHandler) FindByType(ctx context.Context, cmd FindByType) ([]domain.Notification, error) {
	return h.notification.FindByType(ctx, cmd.Type)
}
