package controllers

import (
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func VerificationCodeValidation(router *gin.Engine) {
	router.POST("/verification-code", services.VerifyCode)
}

func ResendVerificationCode(router *gin.Engine) {
	router.GET("/resend-verification-code/:id", services.ResendVerificationCode)
}

func GetVerificationCode(router *gin.Engine) {
	router.GET("/verification-code/:id", services.GetVerificationCode)
}
