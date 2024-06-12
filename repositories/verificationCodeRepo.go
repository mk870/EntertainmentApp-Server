package repositories

import (
	"movieplusApi/db"
	"movieplusApi/models"
)

func CreateVerificationCode(user *models.User, verificationCode models.VerificationCode) bool {
	db.DB.Model(user).Association("RegistrationCode").Replace(verificationCode)
	return true
}

func GetVerificationCodeById(userId string) models.VerificationCode {
	var user models.User
	err := db.DB.Preload("RegistrationCode").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.RegistrationCode
}
