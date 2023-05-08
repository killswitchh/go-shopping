package app

import (
	"go-order-service/middlewares"
	"go-order-service/routers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupApp() *gin.Engine {
	log.Info().Msg("Initializing service")
	app := gin.New()
	app.Use(gin.Recovery())
	app.SetTrustedProxies(nil)
	log.Info().Msg("Adding cors, request id and request logging middleware")
	app.Use(middlewares.CORSMiddleware(), middlewares.RequestID(), middlewares.RequestLogger())
	log.Info().Msg("Setting up routers")
	routers.SetupRouters(app)
	return app

}
