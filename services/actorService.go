package services

import (
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateActorService(c *gin.Context) {
	var actor models.Actor
	err := c.BindJSON(&actor)
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
	isActorCreated := repositories.CreateActorRepository(user, &actor)
	if !isActorCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the actor profile",
		})
		return
	}
	c.String(http.StatusOK, "actor profile has been successfully created")
}

func GetActorsService(c *gin.Context) {
	loggedStudent := c.MustGet("user").(*models.User)
	email := loggedStudent.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	actorList := repositories.GetActorsRepository(user.Id)
	c.JSON(http.StatusOK, actorList)
}

func DeleteActorService(c *gin.Context) {
	actor_id := c.Param("id")
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	actor := repositories.GetActorRepository(user.Id, actor_id)
	isDeleted := repositories.DeleteActorByIdRepository(user, actor)
	if isDeleted {
		c.JSON(http.StatusOK, "actor profile successfully deleted")
		return
	}
}
