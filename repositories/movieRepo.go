package repositories

import (
	"errors"

	"movieplusApi/db"
	"movieplusApi/models"

	"gorm.io/gorm"
)

func CreateMovie(user *models.User, movie *models.Movie) bool {
	err := db.DB.Model(user).Association("Movies").Append(movie)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetMovies(id int) []models.Movie {
	var user = models.User{}
	err := db.DB.Preload("Movies").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Movies
}

func GetMovie(userId int, movieId string) models.Movie {
	var user = models.User{}
	err := db.DB.Preload("Movies", "id=?", movieId).First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Movies[0]
}

func GetUserWithMoviesById(userId int) *models.User {
	var user = models.User{}
	result := db.DB.Preload("Movies").First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func UpdateMovie(user *models.User, updateList []models.Movie) bool {
	user.Movies = updateList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}

func DeleteMovieById(user *models.User, movie models.Movie) bool {
	db.DB.Model(&user).Unscoped().Association("Movies").Delete(movie)
	return true
}
