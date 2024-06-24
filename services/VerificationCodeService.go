package services

import (
	"net/http"
	"strconv"
	"time"

	"movieplusApi/models"
	"movieplusApi/repositories"
	"movieplusApi/tokens"
	"movieplusApi/utilities"

	"github.com/gin-gonic/gin"
)

func CreateVerificationCode() models.VerificationCode {
	var verificationCode models.VerificationCode
	code := utilities.GenerateVerificationCode()
	verificationCode.ExpiryDate = time.Now().Add(time.Minute * 15)
	verificationCode.Code = code
	return verificationCode
}

func VerifyCode(c *gin.Context) {
	type RegistrationCodeDTO struct {
		Id               int
		VerificationCode *int
	}
	var registrationData = RegistrationCodeDTO{}
	c.BindJSON(&registrationData)
	if registrationData.VerificationCode == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter verification code"})
		return
	}
	storedVerificationCode := repositories.GetVerificationCodeById(utilities.ConvertIntToString(registrationData.Id))

	if storedVerificationCode.ExpiryDate.Unix() < time.Now().Local().Unix() {
		isUserDeleted := repositories.DeleteUserById(strconv.Itoa(storedVerificationCode.UserId))
		if !isUserDeleted {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "this user does not exist",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "this verification code has expired please signup again",
			})
			return
		}
	}
	user := repositories.GetUserById(strconv.Itoa(storedVerificationCode.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	refreshToken := tokens.GenerateRefreshToken(user.FirstName, user.Email)
	if refreshToken == "failed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not generate your refresh token",
		})
		return
	}
	user.RefreshToken = refreshToken
	user.IsActive = true
	isUpdated := repositories.SaveUserUpdate(user)
	if isUpdated {
		accessToken := tokens.GenerateAccessToken(user.FirstName, user.Email, user.Id)
		if accessToken == "failed" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not generate your access token",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"accessToken": accessToken,
		})
		return
	}
}

func GetVerificationCode(c *gin.Context) {
	userId := c.Param("id")
	verificationCode := repositories.GetVerificationCodeById(userId)
	c.JSON(http.StatusOK, gin.H{
		"verificationCode": verificationCode,
	})
}

func ResendVerificationCode(c *gin.Context) {
	userId := c.Param("id")
	user := repositories.GetUserById(userId)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	verificationCode := repositories.GetVerificationCodeById(userId)
	verificationCode.Code = utilities.GenerateVerificationCode()
	verificationCode.ExpiryDate = time.Now().Add(time.Minute * 15)
	isVerificationCodeUpdated := repositories.UpdateVerificationCode(&verificationCode)
	if !isVerificationCodeUpdated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate another code",
		})
		return
	}
	isEmailSent := SendVerificationCodeEmail(user.Email, user.FirstName, utilities.ConvertIntToString(verificationCode.Code))
	if !isEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	c.String(http.StatusOK, "please check your email for verification code")
}
