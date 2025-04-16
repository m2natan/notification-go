package queries

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
)

type (
	FindById struct {
		Id string
	}

	FindByIdHandler struct {
		notification domain.NotificationRepository
	}
)

func NewFindByIdHandler(notification domain.NotificationRepository) FindByIdHandler {
	return FindByIdHandler{notification: notification}
}

func (h FindByIdHandler) FindById(ctx context.Context, cmd FindById) (*domain.Notification, error) {
	return h.notification.FindById(ctx, cmd.Id)
}
