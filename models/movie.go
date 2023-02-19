package models

type Movie struct {
	MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	UserId       int    `json:"userId"`
	Title        string `json:"title"`
	Release_date string `json:"release_date"`
	Runtime      string `json:"runtime"`
	Tmdb_id      string `json:"tmdb_id"`
}
