package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateTvShow(router *gin.Engine) {
	router.POST("/tvShow", middleware.AuthValidator, services.CreateTvShow)
}

func GetTvShows(router *gin.Engine) {
	router.GET("/tvShows", middleware.AuthValidator, services.GetTvShows)
}

func DeleteTvShow(router *gin.Engine) {
	router.DELETE("/tvShow/:id", middleware.AuthValidator, services.DeleteTvShow)
}
