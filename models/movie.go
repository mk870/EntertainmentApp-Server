package models

type Movie struct {
	MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	UserId       int    `json:"userId"`
	Title        string `json:"title"`
	Release_date string `json:"release_date"`
	Rating       int    `json:"rating"`
	Tmdb_id      int    `json:"tmdb_id"`
	Poster       string `json:"poster" gorm:"nullable"`
}
