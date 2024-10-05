package db

import (
	"log"
	"os"

	"barr.io/graph/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}

	dsn := "host=localhost user=scoot password=password dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&models.User{})

	db = DB
}
