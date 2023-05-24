package models

type Artist struct {
	MyModel
	Id         int    `json:"id" gorm:"primary_key"`
	UserId     int    `json:"userId"`
	Name       string `json:"name"`
	Followers  string `json:"followers"`
	Spotify_id string `json:"spotify_id"`
	Poster     string `json:"poster" gorm:"nullable"`
}
