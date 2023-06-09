package models

type Actor struct {
	MyModel
	Id         int    `json:"id" gorm:"primary_key"`
	UserId     int    `json:"userId"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Poster     string `json:"poster" gorm:"nullable"`
	Tmdb_id    int    `json:"tmdb_id"`
	Character  string `json:"character"`
}
