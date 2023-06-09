package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
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
	isMovieCreated := repositories.CreateMovie(user, &movie)
	if !isMovieCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the movie",
		})
		return
	}
	c.String(http.StatusOK, "successfully added the movie to your movie list")
}

func GetMovies(c *gin.Context) {
	loggedStudent := c.MustGet("user").(*models.User)
	email := loggedStudent.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	movieList := repositories.GetMovies(user.Id)
	c.JSON(http.StatusOK, movieList)
}

func GetMovie(c *gin.Context) {
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
	movie := repositories.GetMovie(user.Id, movie_id)
	c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
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
	movie := repositories.GetMovie(user.Id, movie_id)
	isDeleted := repositories.DeleteMovieById(user, movie)
	if isDeleted {
		c.JSON(http.StatusOK, "movie successfully deleted")
		return
	}
}
