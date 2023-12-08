package helpers

import (
	"fmt"
	"go-notification-service/hooks"

	"gorm.io/gorm"
)

func Bridge(tx *gorm.DB) {
	fmt.Print("helper")
	hooks.GetLatestNotification(tx)
}