package repositories

import (
	"errors"

	"movieplusApi/db"
	"movieplusApi/models"

	"gorm.io/gorm"
)

func CreateAlbum(user *models.User, album *models.Album) bool {
	err := db.DB.Model(user).Association("Albums").Append(album)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetAlbums(id int) []models.Album {
	var user = models.User{}
	err := db.DB.Preload("Albums").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Albums
}

func GetAlbum(userId int, albumId string) models.Album {
	var user = models.User{}
	err := db.DB.Preload("Albums", "id=?", albumId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Albums[0]
}

func GetUserWithAlbumsById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Albums").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateAlbum(user *models.User, updateList []models.Album) bool {
	user.Albums = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteAlbumById(user *models.User, album models.Album) bool {
	db.DB.Model(&user).Unscoped().Association("Albums").Delete(album)
	return true
}
