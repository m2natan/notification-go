package domain

import (
	"errors"
	"log"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/notifier"
	"gorm.io/gorm"
)

type EmailStatus string
type NotificationType string

var ErrNotificationIsBlank = errors.New("notification field is blank")

const (
	StatusPending EmailStatus = "pending"
	StatusSent    EmailStatus = "sent"
	StatusFailed  EmailStatus = "failed"

	NotificationTypeEmail   NotificationType = "email"
	NotificationTypeSms     NotificationType = "sms"
	NotificationTypeUnknown NotificationType = "unknown"
)

type Notification struct {
	gorm.Model
	Id            string           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Subject       string           `json:"subject"`
	Content       string           `json:"content"`
	SenderName    string           `json:"sender_name"`
	Sender        string           `json:"sender"`
	Recipient     string           `json:"recipient"`
	RecipientName string           `json:"recipient_name"`
	Status        EmailStatus      `json:"status"`
	Type          NotificationType `json:"type"`
}

func CreateNotification(subject string, content string, senderName string, sender string, recipient string, recipientName string, status EmailStatus, notification_type NotificationType) (notification *Notification, err error) {
	// Input validation
	if subject == "" || content == "" || senderName == "" || sender == "" || recipient == "" || recipientName == "" || notification_type == "" {
		return nil, ErrNotificationIsBlank
	}

	// Validate and map the string type to NotificationType enum
	var notifType NotificationType
	switch notification_type {
	case "email":
		notifType = NotificationTypeEmail
	case "sms":
		notifType = NotificationTypeSms
	default:
		return nil, errors.New("invalid notification type") // Error handling for unsupported type
	}

	// Create the notification
	notification = &Notification{
		Subject:       subject,
		Content:       content,
		SenderName:    senderName,
		Sender:        sender,
		Recipient:     recipient,
		RecipientName: recipientName,
		Status:        StatusPending,
		Type:          notifType,
	}

	// Start a goroutine to send the notification asynchronously without blocking the rest of the code
	go func() {
		err := CreateAndSendNotification(notification)
		if err != nil {
			log.Printf("Error sending notification: %v", err) // Logging the error, or you could use a monitoring service
		}
	}()

	// Return the created notification immediately without waiting for the goroutine
	return notification, nil
}

// CreateAndSendNotification creates the notification and sends it using the provided notifier
func CreateAndSendNotification(cmd *Notification) error {
	// Send the notification (this could be an email or SMS depending on cmd.Type)
	err := notifier.Send(cmd.Subject, cmd.Content, cmd.SenderName, cmd.Sender, cmd.Recipient, cmd.RecipientName, string(cmd.Type))
	return err
}
