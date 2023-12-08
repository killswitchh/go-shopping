package hooks

import (
	"fmt"
	"go-notification-service/models"

	"gorm.io/gorm"
)

func AfterCreateCallback(tx *gorm.DB){
	fmt.Println("after_create_callback")
	GetLatestNotification(tx)
}

func GetLatestNotification(tx *gorm.DB) {
	fmt.Println("inside")
	var latestNotification models.Notification
	if err := tx.Last(&latestNotification).Error; err != nil {
		fmt.Println("Error fetching latest record:")
		return
	}

	fmt.Printf("Latest record: ID=%d, Title=%s, Body=%s\n", latestNotification.OrderId, latestNotification.Order)
}