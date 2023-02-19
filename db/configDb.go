package db

import (
	"movieplusApi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://cscujoix:HeSrlNYoJmFHZTzN-6-0u0nFCeJIXDl8@trumpet.db.elephantsql.com/cscujoix"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Actor{}, &models.VerificationToken{}, &models.Movie{})
	DB = db
}
