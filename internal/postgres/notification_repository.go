package postgres

import (
	"context"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

var _ domain.NotificationRepository = (*NotificationRepository)(nil)

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// Create implements domain.NotificationRepository.
func (n *NotificationRepository) Create(ctx context.Context, notification *domain.Notification) error {
	return n.db.WithContext(ctx).Create(notification).Error
}

// Delete implements domain.NotificationRepository.
func (n *NotificationRepository) Delete(ctx context.Context, Id string) error {
	return n.db.WithContext(ctx).Where("id = ?", Id).Delete(&domain.Notification{}).Error
}

// FindAll implements domain.NotificationRepository.
func (n *NotificationRepository) FindAll(ctx context.Context) ([]domain.Notification, error) {
	var notification []domain.Notification
	if err := n.db.WithContext(ctx).Find(&notification).Error; err != nil {
		return nil, err
	}
	return notification, nil

}

// FindById implements domain.NotificationRepository.
func (n *NotificationRepository) FindById(ctx context.Context, ID string) (*domain.Notification, error) {
	var notification domain.Notification
	if err := n.db.WithContext(ctx).Where("id = ?", ID).First(&notification).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

// FindByStatus implements domain.NotificationRepository.
func (n *NotificationRepository) FindByStatus(ctx context.Context, status domain.EmailStatus) ([]domain.Notification, error) {
	var notification []domain.Notification
	if err := n.db.WithContext(ctx).Where("status = ?", status).Find(&notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

// FindByType implements domain.NotificationRepository.
func (n *NotificationRepository) FindByType(ctx context.Context, notification_type domain.NotificationType) ([]domain.Notification, error) {
	var notification []domain.Notification
	if err := n.db.WithContext(ctx).Where("type = ?", notification_type).Find(&notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

// Update implements domain.NotificationRepository.
func (n *NotificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	return n.db.WithContext(ctx).Where("id = ?", notification.ID).Updates(notification).Error
}
