package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateMovieService(c *gin.Context) {
	var movie models.Movie
	err := c.BindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not bind the request body",
		})
		return
	}
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find this user",
		})
		return
	}
	isMovieCreated := repositories.CreateMovieRepository(user, &movie)
	if !isMovieCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the movie",
		})
		return
	}
	c.String(http.StatusOK, "successfully added the movie to your movie list")
}

func GetMoviesService(c *gin.Context) {
	loggedStudent := c.MustGet("user").(*models.User)
	email := loggedStudent.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	movieList := repositories.GetMoviesRepository(user.Id)
	c.JSON(http.StatusOK, movieList)
}

func GetMovieService(c *gin.Context) {
	movie_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	movie := repositories.GetMovieRepository(user.Id, movie_id)
	c.JSON(http.StatusOK, movie)
}

func DeleteMovieService(c *gin.Context) {
	movie_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	movie := repositories.GetMovieRepository(user.Id, movie_id)
	isDeleted := repositories.DeleteMovieByIdRepository(user, movie)
	if isDeleted {
		c.JSON(http.StatusOK, "movie successfully deleted")
		return
	}
}
