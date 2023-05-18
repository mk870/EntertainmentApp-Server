package db

import (
	"movieplusApi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:mkhue@localhost:5420/students"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Actor{}, &models.VerificationToken{}, &models.Movie{})
	DB = db
}
