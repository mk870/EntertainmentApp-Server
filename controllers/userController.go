package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.Engine) {
	router.POST("/user", services.CreateUserService)
}

func GetUsers(router *gin.Engine) {
	router.GET("/users", middleware.AuthValidator, services.GetUsersService)
}

func UpdateUser(router *gin.Engine) {
	router.PUT("/user/:id", middleware.AuthValidator, services.UpdateUserService)
}

func GetUser(router *gin.Engine) {
	router.GET("/user/:id", middleware.AuthValidator, services.GetUserService)
}

func DeleteUser(router *gin.Engine) {
	router.DELETE("/user/:id", middleware.AuthValidator, services.DeleteUserService)
}
