package repositories

import (
	"errors"

	"movieplusApi/db"
	"movieplusApi/models"

	"gorm.io/gorm"
)

func CreateActor(user *models.User, actor *models.Actor) bool {
	err := db.DB.Model(user).Association("Actors").Append(actor)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetActors(id int) []models.Actor {
	var user = models.User{}
	err := db.DB.Preload("Actors").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Actors
}

func GetActor(userId int, actorId string) models.Actor {
	var user = models.User{}
	err := db.DB.Preload("Actors", "id=?", actorId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Actors[0]
}

func GetUserWithActorsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Actors").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateActor(user *models.User, updateList []models.Actor) bool {
	user.Actors = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteActorById(user *models.User, actor models.Actor) bool {
	db.DB.Model(&user).Unscoped().Association("Actors").Delete(actor)
	return true
}
