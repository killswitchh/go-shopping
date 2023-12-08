package main

import (
	"go-notification-service/consts"
	"go-notification-service/database"
	"go-notification-service/handlers"
	"go-notification-service/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")
	database.ConnectDb()
	exampleQueue := utils.RMQConsumer{
		Queue: consts.QUEUE,
		ConnectionString: connectionString,
		MsgHandler: handlers.ConsumeMessage,
	}
	forever := make(chan bool)
	go exampleQueue.Consume()
	<-forever
}
