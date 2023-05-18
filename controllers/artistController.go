package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateArtist(router *gin.Engine) {
	router.POST("/artist", middleware.AuthValidator, services.CreateArtist)
}

func GetArtists(router *gin.Engine) {
	router.GET("/artists", middleware.AuthValidator, services.GetArtists)
}

func DeleteArtist(router *gin.Engine) {
	router.DELETE("/artist/:id", middleware.AuthValidator, services.DeleteArtist)
}
