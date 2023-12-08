package models

import (
	"go-notification-service/interfaces"
	"time"

	"gorm.io/gorm"
)
type Notification struct {
	OrderId          uint
	Order            Order
	NotificationType string
	Recipient        string
	Status           string `gorm:"default:pending"`
	CreatedAt        time.Time
}

func (u *Notification) AfterCreate(tx *gorm.DB) (err error) {
    HelperNotification(tx)
	return
}

func HelperNotification(tx *gorm.DB) {
	var o interfaces = interfaces.NotificationInterface{}
	o.GetLatestNotification(tx)
}