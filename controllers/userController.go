package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.Engine) {
	router.POST("/user", services.CreateUser)
}
func CreateUserByMobile(router *gin.Engine) {
	router.POST("/user-mobile", services.UserRegistrationByMobile)
}

func GetUsers(router *gin.Engine) {
	router.GET("/users", middleware.AuthValidator, services.GetUsers)
}

func UpdateUser(router *gin.Engine) {
	router.PUT("/user/:id", middleware.AuthValidator, services.UpdateUser)
}

func GetUser(router *gin.Engine) {
	router.GET("/user/:id", middleware.AuthValidator, services.GetUser)
}

func DeleteUser(router *gin.Engine) {
	router.DELETE("/user/:id", middleware.AuthValidator, services.DeleteUser)
}
