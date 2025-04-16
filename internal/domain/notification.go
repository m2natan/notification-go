package domain

import (
	"errors"

	"gorm.io/gorm"
)

type EmailStatus string
type NotifiationType string

var ErrNotificationIsBlank = errors.New("notification filed is blank")

const (
	StatusPending EmailStatus = "pending"
	StatusSent    EmailStatus = "sent"
	StatusFailed  EmailStatus = "failed"

	NotificationTypeEmail NotifiationType = "email"
	NotificationTypeSms   NotifiationType = "sms"
)

type Notification struct {
	gorm.Model
	Id             string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Subject        string
	Content        string
	SenderName     string
	SenderEmail    string
	RecipientEmail string
	RecipientName  string
	Status         EmailStatus
	Type           NotifiationType
}

func CreateNotification(subject string, content string, senderName string, senderEmail string, recipientEmail string, recipientName string, status EmailStatus, notification_type NotifiationType) (notification *Notification, err error) {
	if subject == "" || content == "" || senderName == "" || senderEmail == "" || recipientEmail == "" || recipientName == "" || notification_type == "" {
		return nil, ErrNotificationIsBlank
	}

	notification = &Notification{
		Subject:        subject,
		Content:        content,
		SenderName:     senderName,
		SenderEmail:    senderEmail,
		RecipientEmail: recipientEmail,
		RecipientName:  recipientName,
		Status:         status,
		Type:           notification_type,
	}
	return
}
