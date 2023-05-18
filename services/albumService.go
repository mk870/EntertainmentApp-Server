package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateAlbum(c *gin.Context) {
	var album models.Album
	err := c.BindJSON(&album)
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
	isAlbumCreated := repositories.CreateAlbum(user, &album)
	if !isAlbumCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the album",
		})
		return
	}
	c.String(http.StatusOK, "album has been successfully created")
}

func GetAlbums(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	albumList := repositories.GetAlbums(user.Id)
	c.JSON(http.StatusOK, albumList)
}

func DeleteAlbum(c *gin.Context) {
	album_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	album := repositories.GetAlbum(user.Id, album_id)
	isDeleted := repositories.DeleteAlbumById(user, album)
	if isDeleted {
		c.JSON(http.StatusOK, "album successfully deleted")
		return
	}
}
