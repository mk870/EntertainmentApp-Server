package db

import (
	"log"
	"movieplusApi/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dns := os.Getenv("DATABASE_DETAILS")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Actor{}, &models.VerificationToken{}, &models.Movie{}, &models.Album{}, &models.Artist{}, &models.Track{}, &models.TvShow{})
	DB = db
}
