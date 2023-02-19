package controllers

import (
	"movieplusApi/middleware"
	"movieplusApi/services"

	"github.com/gin-gonic/gin"
)

func CreateMovie(router *gin.Engine) {
	router.POST("/movie", middleware.AuthValidator, services.CreateMovieService)
}

func GetMovies(router *gin.Engine) {
	router.GET("/movies", middleware.AuthValidator, services.GetMoviesService)
}

func DeleteMovie(router *gin.Engine) {
	router.DELETE("/movie/:id", middleware.AuthValidator, services.DeleteMovieService)
}
