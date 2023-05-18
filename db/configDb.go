package db

import (
	"movieplusApi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dns := "postgres://izesjfds:jKWVS1IZHYbArtwFxfbiAlb3uHbLxfVp@horton.db.elephantsql.com/izesjfds"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Actor{}, &models.VerificationToken{}, &models.Movie{})
	DB = db
}
