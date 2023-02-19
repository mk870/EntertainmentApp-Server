package models

type Actor struct {
	MyModel
	Id          int    `json:"id" gorm:"primary_key"`
	UserId      int    `json:"userId"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	Birth_place string `json:"birth_place"`
	Age         string `json:"age"`
	Tmdb_id     string `json:"tmdb_id"`
}
