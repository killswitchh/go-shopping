package interfaces

import (
	"gorm.io/gorm"
)

type NotificationInterface interface {
	GetLatestNotification(tx *gorm.DB)
}