package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateTrack(c *gin.Context) {
	var track models.Track
	err := c.BindJSON(&track)
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
	isTrackCreated := repositories.CreateTrack(user, &track)
	if !isTrackCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the track",
		})
		return
	}
	c.String(http.StatusOK, "the track has been successfully created")
}

func GetTracks(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	trackList := repositories.GetTracks(user.Id)
	c.JSON(http.StatusOK, trackList)
}

func DeleteTrack(c *gin.Context) {
	track_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	track := repositories.GetTrack(user.Id, track_id)
	isDeleted := repositories.DeleteTrackById(user, track)
	if isDeleted {
		c.JSON(http.StatusOK, "the track was successfully deleted")
		return
	}
}
