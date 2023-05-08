package main

import (
	"go-notification-service/consts"
	"go-notification-service/handlers"
	"go-notification-service/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")
	exampleQueue := utils.RMQConsumer{
		consts.EXAMPLE_QUEUE,
		connectionString,
		handlers.HandleExample,
	}
	forever := make(chan bool)
	go exampleQueue.Consume()
	<-forever
}
