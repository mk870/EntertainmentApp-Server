package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateActor(router *gin.Engine) {
	router.POST("/actor", middleware.AuthValidator, services.CreateActor)
}

func GetActors(router *gin.Engine) {
	router.GET("/actors", middleware.AuthValidator, services.GetActors)
}

func DeleteActor(router *gin.Engine) {
	router.DELETE("/actor/:id", middleware.AuthValidator, services.DeleteActor)
}
