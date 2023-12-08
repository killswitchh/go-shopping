package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"

	"go-notification-service/models"
)

func ConsumeMessage(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		log.Err(err).Msg("Error occurred in RMQ consumer")
	}
	log.Info().Msgf("Message received on '%s' queue: %s", queue, string(msg.Body))

	order := models.Order{}
	res := json.Unmarshal(msg.Body, &order)
    if res != nil {
		fmt.Print(res)
        panic(res)
    }
	StoreNotification(order)
	// SendMail(order)
}