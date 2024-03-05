package controllers

import (
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func GetNews(router *gin.Engine) {
	router.GET("/news/:category", services.GetNews)
}
