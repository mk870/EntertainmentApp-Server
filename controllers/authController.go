package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func GetNewAccessToken(router *gin.Engine) {
	router.GET("/refresh-token", services.RefreshAccessTokenService)
}

func Login(router *gin.Engine) {
	router.POST("/login", services.LoginService)
}

func LoginOut(router *gin.Engine) {
	router.GET("/logout", middleware.AuthValidator, services.LogoutService)
}
