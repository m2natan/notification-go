package notifier

import (
	"fmt"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/notifier/email"
)

func Send(subject, content, senderName, sender, recipient, recipientName, Type string) error {
	if Type == "email" {
		return email.NewMailjetNotifier().Send(subject, content, senderName, sender, recipient, recipientName)
	} else if Type == "sms" {
		return email.NewMailjetNotifier().Sms(recipient, content)
	} else {
		return fmt.Errorf("unknown notification type: %s", Type)
	}
}
