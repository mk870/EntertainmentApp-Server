package repositories

import (
	"errors"

	"movieplusApi/db"
	"movieplusApi/models"

	"gorm.io/gorm"
)

func CreateTrack(user *models.User, track *models.Track) bool {
	err := db.DB.Model(user).Association("Tracks").Append(track)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetTracks(id int) []models.Track {
	var user = models.User{}
	err := db.DB.Preload("Tracks").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Tracks
}

func GetTrack(userId int, trackId string) models.Track {
	var user = models.User{}
	err := db.DB.Preload("Tracks", "id=?", trackId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Tracks[0]
}

func GetUserWithTracksById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Tracks").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateTrack(user *models.User, updateList []models.Track) bool {
	user.Tracks = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteTrackById(user *models.User, track models.Track) bool {
	db.DB.Model(&user).Unscoped().Association("Tracks").Delete(track)
	return true
}
