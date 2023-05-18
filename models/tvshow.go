package models

type TvShow struct {
	MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	UserId       int    `json:"userId"`
	Name         string `json:"name"`
	Release_date string `json:"release_date"`
	Rating       string `json:"rating"`
	Tmdb_id      int    `json:"tmdb_id"`
	Poster       string `json:"poster" gorm:"nullable"`
}
