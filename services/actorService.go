package services

import (
	"fmt"
	"net/http"

	"movieplusApi/models"
	"movieplusApi/repositories"

	"github.com/gin-gonic/gin"
)

func CreateActor(c *gin.Context) {
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
	isActorCreated := repositories.CreateActor(user, &actor)
	if !isActorCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create the actor profile",
		})
		return
	}
	c.String(http.StatusOK, "actor profile has been successfully created")
}

func GetActors(c *gin.Context) {
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	actorList := repositories.GetActors(user.Id)
	c.JSON(http.StatusOK, actorList)
}

func DeleteActor(c *gin.Context) {
	actor_id := c.Param("id")
	fmt.Println("actorid", actor_id)
	loggedInUser := c.MustGet("user").(*models.User)
	email := loggedInUser.Email
	user := repositories.GetUserByEmail(email)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not find the user",
		})
		return
	}
	actor := repositories.GetActor(user.Id, actor_id)
	isDeleted := repositories.DeleteActorById(user, actor)
	if isDeleted {
		c.JSON(http.StatusOK, "actor profile successfully deleted")
		return
	}
}
