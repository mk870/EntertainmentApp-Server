package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func GetNewAccessToken(router *gin.Engine) {
	router.GET("/refresh-token", services.RefreshAccessToken)
}

func Login(router *gin.Engine) {
	router.POST("/login", services.Login)
}

func LoginOut(router *gin.Engine) {
	router.GET("/logout", middleware.AuthValidator, services.Logout)
}
