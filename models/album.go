package models

type Album struct {
	MyModel
	Id         int    `json:"id" gorm:"primary_key"`
	UserId     int    `json:"userId"`
	Name       string `json:"name"`
	Artists    string `json:"artists"`
	Spotify_id string `json:"spotify_id"`
	Poster     string `json:"poster" gorm:"nullable"`
}
