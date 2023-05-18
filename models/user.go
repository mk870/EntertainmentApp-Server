package models

import (
	"time"

	"gorm.io/gorm"
)

type MyModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	MyModel
	Id                int               `json:"id" gorm:"primary_key"`
	FirstName         string            `json:"firstName" validate:"required,min=2,max=50"`
	LastName          string            `json:"lastName" validate:"required,min=2,max=50"`
	Email             string            `json:"email" gorm:"unique" validate:"email,required"`
	Password          string            `json:"password" validate:"required,min=2,max=50"`
	RefreshToken      string            `json:"refreshToken"`
	IsActive          bool              `json:"isActive"`
	Movies            []Movie           `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Actors            []Actor           `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tracks            []Track           `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TvShows           []TvShow          `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Albums            []Album           `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Artists           []Artist          `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RegistrationToken VerificationToken `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
