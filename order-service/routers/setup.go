package routers

import (
	"go-order-service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		orders := v1.Group("/orders")
		{
			orders.POST("/create", controllers.CreateOrder)
		}
		v1.GET("/ping", controllers.Ping)
		v1.GET("/ping2", controllers.Ping)
		v1.POST("/publish/example", controllers.Example)
	}
}
