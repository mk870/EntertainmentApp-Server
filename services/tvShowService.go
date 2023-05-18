package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateTvShow(c *gin.Context) {
	var tvShow models.TvShow
	err := c.BindJSON(&tvShow)
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
	isTvShowCreated := repositories.CreateTvShow(user, &tvShow)
	if !isTvShowCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the tvShow",
		})
		return
	}
	c.String(http.StatusOK, "the tvShow has been successfully created")
}

func GetTvShows(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	tvShowList := repositories.GetTvShows(user.Id)
	c.JSON(http.StatusOK, tvShowList)
}

func DeleteTvShow(c *gin.Context) {
	tvShow_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	tvShow := repositories.GetTvShow(user.Id, tvShow_id)
	isDeleted := repositories.DeleteTvShowById(user, tvShow)
	if isDeleted {
		c.JSON(http.StatusOK, "the tvShow was successfully deleted")
		return
	}
}
