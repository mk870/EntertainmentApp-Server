package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateArtist(c *gin.Context) {
	var artist models.Artist
	err := c.BindJSON(&artist)
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
	isArtistCreated := repositories.CreateArtist(user, &artist)
	if !isArtistCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the artist profile",
		})
		return
	}
	c.String(http.StatusOK, "artist profile has been successfully created")
}

func GetArtists(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	artistList := repositories.GetArtists(user.Id)
	c.JSON(http.StatusOK, artistList)
}

func DeleteArtist(c *gin.Context) {
	artist_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	artist := repositories.GetArtist(user.Id, artist_id)
	isDeleted := repositories.DeleteArtistById(user, artist)
	if isDeleted {
		c.JSON(http.StatusOK, "artist profile successfully deleted")
		return
	}
}
