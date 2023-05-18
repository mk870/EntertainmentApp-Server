package repositories

import (
	"errors"

	"movieplusApi/db"
	"movieplusApi/models"

	"gorm.io/gorm"
)

func CreateArtist(user *models.User, artist *models.Artist) bool {
	err := db.DB.Model(user).Association("Artists").Append(artist)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetArtists(id int) []models.Artist {
	var user = models.User{}
	err := db.DB.Preload("Artists").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Artists
}

func GetArtist(userId int, artistId string) models.Artist {
	var user = models.User{}
	err := db.DB.Preload("Artists", "id=?", artistId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Artists[0]
}

func GetUserWithArtistsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Artists").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateArtist(user *models.User, updateList []models.Artist) bool {
	user.Artists = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteArtistById(user *models.User, artist models.Artist) bool {
	db.DB.Model(&user).Unscoped().Association("Artists").Delete(artist)
	return true
}
