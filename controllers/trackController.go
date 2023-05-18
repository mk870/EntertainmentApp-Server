package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateTrack(router *gin.Engine) {
	router.POST("/track", middleware.AuthValidator, services.CreateTrack)
}

func GetTracks(router *gin.Engine) {
	router.GET("/tracks", middleware.AuthValidator, services.GetTracks)
}

func DeleteTrack(router *gin.Engine) {
	router.DELETE("/track/:id", middleware.AuthValidator, services.DeleteTrack)
}
