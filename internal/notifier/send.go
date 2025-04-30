package notifier

import (
	"fmt"

	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/notifier/email"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/notifier/push"
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/notifier/sms"
)

func Send(subject, content, senderName, sender, recipient, recipientName, Type string) error {
	if Type == "email" {
		return email.NewMailjetNotifier().Send(subject, content, senderName, sender, recipient, recipientName)
	} else if Type == "sms" {
		return sms.NewTwilioNotifier().SendSms(recipient, content)
	} else if Type == "push"{
		return push.SendNotification(recipient, content)
	}else {
		return fmt.Errorf("unknown notification type: %s", Type)
	}
}
