package controllers

import (
	"go-order-service/consts"
	"go-order-service/models"
	"go-order-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Example(c *gin.Context) {
	var msg models.Message

	request_id := c.GetString("x-request-id")

	if binderr := c.ShouldBindJSON(&msg); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	connectionString := utils.GetEnvVar("RMQ_URL")

	rmqProducer := utils.RMQProducer{
		consts.EXAMPLE_QUEUE,
		connectionString,
	}

	rmqProducer.PublishMessage("text/plain", []byte(msg.Message))

	c.JSON(http.StatusOK, gin.H{
		"response": "Message received",
	})

}
