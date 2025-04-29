package domain

import "context"

type NotificationRepository interface {
	Create(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
	Delete(ctx context.Context, Id string) error
	FindById(ctx context.Context, Id string) (*Notification, error)
	FindAll(ctx context.Context) ([]Notification, error)
	FindByStatus(ctx context.Context, status EmailStatus) ([]Notification, error)
	FindByType(ctx context.Context, notification_type NotificationType) ([]Notification, error)
}
