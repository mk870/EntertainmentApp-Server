package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateAlbum(router *gin.Engine) {
	router.POST("/album", middleware.AuthValidator, services.CreateAlbum)
}

func GetAlbums(router *gin.Engine) {
	router.GET("/albums", middleware.AuthValidator, services.GetAlbums)
}

func DeleteAlbum(router *gin.Engine) {
	router.DELETE("/album/:id", middleware.AuthValidator, services.DeleteAlbum)
}
