package repositories

import (
	"errors"

	"movieplusApi/db"
	"movieplusApi/models"

	"gorm.io/gorm"
)

func CreateTvShow(user *models.User, tvShow *models.TvShow) bool {
	err := db.DB.Model(user).Association("TvShows").Append(tvShow)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetTvShows(id int) []models.TvShow {
	var user = models.User{}
	err := db.DB.Preload("TvShows").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.TvShows
}

func GetTvShow(userId int, tvShowId string) models.TvShow {
	var user = models.User{}
	err := db.DB.Preload("TvShows", "id=?", tvShowId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.TvShows[0]
}

func GetUserWithTvShowsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("TvShows").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateTvShow(user *models.User, updateList []models.TvShow) bool {
	user.TvShows = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteTvShowById(user *models.User, tvShow models.TvShow) bool {
	db.DB.Model(&user).Unscoped().Association("TvShows").Delete(tvShow)
	return true
}
