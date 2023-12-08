package handlers

import (
	"go-notification-service/database"
	"go-notification-service/models"
)

func StoreNotification(order models.Order) {
	notification := models.Notification{
		Order: order,
		OrderId: order.ID,
		NotificationType: "email",
		Recipient: order.Email,
		Status: "pending",
	}
	result := database.DB.Omit("Order").Create(&notification)
	if result.Error != nil {
		panic(result.Error)
	}
}