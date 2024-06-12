package models

import "time"

type VerificationCode struct {
	MyModel
	Id         int       `json:"id" gorm:"primary_key"`
	UserId     int       `json:"userId"`
	Code       int       `json:"code"`
	ExpiryDate time.Time `json:"expirydate"`
}
