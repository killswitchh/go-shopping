package controllers

import (
	"encoding/json"
	"fmt"
	"go-order-service/consts"
	"go-order-service/database"
	"go-order-service/models"
	"go-order-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	log.Info().Msgf("Creating order :%s", order.ID)
	if err := c.BindJSON(&order); err != nil {
		return
	}
	database.DB.Db.Create(&order)
	PublishOrder(order)
	fmt.Printf("Order created and published: %+v", order)
	c.JSON(http.StatusOK, order)
}

func PublishOrder(msg models.Order) {
	connectionString := utils.GetEnvVar("RMQ_URL")
	rmqProducer := utils.RMQProducer{
		Queue: consts.ORDER_QUEUE,
		ConnectionString: connectionString,
	}
	fmt.Printf("PRODUCER: %+v", rmqProducer)

	// convert msg to json
	jsonify, err := json.Marshal(msg)

	if err != nil {
		return
	}
	fmt.Printf("JSONIFY: %+v", jsonify)
	rmqProducer.PublishMessage("application/json", jsonify)
}
