package email

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go"
)

type MailjetNotifier struct {
	client *mailjet.Client
}

// NewMailjetNotifier creates an instance of MailjetNotifier with API keys
func NewMailjetNotifier() *MailjetNotifier {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	privateKey := os.Getenv("MJ_APIKEY_PRIVATE")
	client := mailjet.NewMailjetClient(publicKey, privateKey)

	return &MailjetNotifier{client: client}
}

func (m *MailjetNotifier) Send(subject string, content string, senderName string, sender string, recipient string, recipientName string) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: sender,
				Name:  senderName,
			},
			To: &mailjet.RecipientsV31{
				{
					Email: recipient,
					Name:  recipientName,
				},
			},
			Subject:  subject,
			HTMLPart: content,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := m.client.SendMailV31(&messages)
	if err != nil {
		return fmt.Errorf("mailjet send error: %w", err)
	}
	return nil
}

