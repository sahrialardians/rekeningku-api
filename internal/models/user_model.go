package models

import "time"

type User struct {
	ID        int       `gorm:"type:int;primary_key"`
	Fullname  string    `gorm:"type:varchar(255)" validate:"required,min=3,max=255"`
	Email     string    `gorm:"type:varchar(255);unique" validate:"required,email"`
	Password  string    `gorm:"type:varchar(255)" validate:"required,min=6"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
