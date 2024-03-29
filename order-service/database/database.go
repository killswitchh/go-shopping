package database

import (
	"fmt"
	"go-order-service/models"
	"go-order-service/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	fmt.Println(utils.GetEnvVar("POSTGRES_HOST"))
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		utils.GetEnvVar("POSTGRES_HOST"),
		utils.GetEnvVar("POSTGRES_USER"),
		utils.GetEnvVar("POSTGRES_PASSWORD"),
		utils.GetEnvVar("POSTGRES_NAME"),
		utils.GetEnvVar("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the Database %s", dsn)
	}
	fmt.Println("? Connected Successfully to the Database")
	if err != nil {
		log.Fatalf("Failed to connect to database.  %s", dsn)
		os.Exit(2)
	}

	log.Printf("connected %s", dsn)
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Order{})

	DB = Dbinstance{
		Db: db,
	}
}
