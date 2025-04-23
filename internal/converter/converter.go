package converter

import (
	"github.com/Kifiya-Financial-Technology/Notification-Service/internal/domain"
	"github.com/Kifiya-Financial-Technology/Notification-Service/notificationpb"
)

func ConvertNotificationToPb(from *domain.Notification) (to *notificationpb.Notification) {
	to = &notificationpb.Notification{
		Id:             from.Id,
		Subject:        from.Subject,
		Content:        from.Content,
		SenderName:     from.SenderName,
		SenderEmail:    from.SenderEmail,
		RecipientEmail: from.RecipientEmail,
		RecipientName:  from.RecipientName,
		Status:         convertStatusToPb(from.Status),
		Type:           convertTypeToPb(from.Type),
	}
	return
}

func convertStatusToPb(status domain.EmailStatus) notificationpb.EmailStatus {
	switch status {
	case domain.StatusPending:
		return notificationpb.EmailStatus_STATUS_PENDING
	case domain.StatusSent:
		return notificationpb.EmailStatus_STATUS_SENT
	default:
		return notificationpb.EmailStatus_STATUS_FAILED
	}
}

func convertTypeToPb(t domain.NotifiationType) notificationpb.NotificationType {
	switch t {
	case domain.NotificationTypeSms:
		return notificationpb.NotificationType_NOTIFICATION_TYPE_SMS
	case domain.NotificationTypeEmail:
		return notificationpb.NotificationType_NOTIFICATION_TYPE_EMAIL
	default:
		return notificationpb.NotificationType_NOTIFICATION_TYPE_UNKNOWN
	}
}

func ConvertPbTypeToDomain(t notificationpb.NotificationType) domain.NotifiationType {
	switch t {
	case notificationpb.NotificationType_NOTIFICATION_TYPE_SMS:
		return domain.NotificationTypeSms
	case notificationpb.NotificationType_NOTIFICATION_TYPE_EMAIL:
		return domain.NotificationTypeEmail
	default:
		return domain.NotificationTypeUnknown
	}
}
