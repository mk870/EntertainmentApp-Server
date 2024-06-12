package db

import (
	"movieplusApi/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dns := os.Getenv("DATABASE_DETAILS")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Actor{}, &models.VerificationToken{}, &models.VerificationCode{}, &models.Movie{}, &models.Album{}, &models.Artist{}, &models.Track{}, &models.TvShow{})
	DB = db
}
