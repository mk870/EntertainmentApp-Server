package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"
	"movieplusApi/tokens"
	"movieplusApi/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistrationByMobile(c *gin.Context) {
	var user models.User
	validateModelFields := validator.New()
	c.BindJSON(&user)

	modelFieldsValidationError := validateModelFields.Struct(user)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	refreshTokenResult := tokens.GenerateRefreshToken(user.FirstName, user.Email)
	if refreshTokenResult == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create refresh token",
		})
		return
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(hashedPassword),
		Email:     user.Email,
		RegistrationCode: models.VerificationCode{
			Code:       utilities.GenerateVerificationCode(),
			ExpiryDate: utilities.GenerateVerificationGracePeriod(),
		},
	}

	isUserCreated := repositories.CreateUser(&newUser)
	if !isUserCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "this email already exists"})
		return
	}

	isVerificationEmailSent := SendVerificationCodeEmail(user.Email, user.FirstName, utilities.ConvertIntToString(newUser.RegistrationCode.Code))
	if !isVerificationEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	createdUser := repositories.GetUserByEmail(user.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "please check your email for verification",
		"userId":  createdUser.Id,
	})

}
