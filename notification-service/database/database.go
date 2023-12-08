package database

import (
	"fmt"

	"go-notification-service/models"
	"go-notification-service/utils"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDb() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

  db_name, db_host, db_port, user, db_pass, ssl_mode, timezone := 
	utils.GetEnvVar("DB_NAME"),
	utils.GetEnvVar("DB_HOST"),
	utils.GetEnvVar("DB_PORT"),
	utils.GetEnvVar("USER"),
	utils.GetEnvVar("DB_PASS"),
	utils.GetEnvVar("SSL_MODE"),
	utils.GetEnvVar("TIMEZONE")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db_host, user, db_pass,db_name, db_port, ssl_mode, timezone)
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
	log.Logger.Fatal().Msgf("Failed to connect to the Database %s", dsn)
  }
  
  log.Logger.Info().Msg("Connected to DB instance")
  DB.AutoMigrate(&models.Notification{})
}