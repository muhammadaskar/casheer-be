package notification

import (
	"time"

	"github.com/muhammadaskar/casheer-be/domains"
)

type NotificationFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      int       `json:"type"`
	UserId    int       `json:"user_id"`
	ProductId int       `json:"product_id"`
	IsRead    int       `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatNotification(notification domains.Notification) NotificationFormatter {
	notificationFormatter := NotificationFormatter{}
	notificationFormatter.ID = notification.ID
	notificationFormatter.Name = notification.Name
	notificationFormatter.Type = notification.Type
	notificationFormatter.UserId = notification.UserId
	notificationFormatter.ProductId = notification.ProductId
	notificationFormatter.IsRead = notification.IsRead
	notificationFormatter.CreatedAt = notification.CreatedAt

	return notificationFormatter
}

func NotificationsFormatter(notifications []domains.Notification) []NotificationFormatter {
	notificationsFormattter := []NotificationFormatter{}
	for _, nonotification := range notifications {
		notificationFormattter := FormatNotification(nonotification)
		notificationsFormattter = append(notificationsFormattter, notificationFormattter)
	}

	return notificationsFormattter
}
