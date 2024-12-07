package models

import "time"

type Account struct {
	ID                int       `gorm:"type:int;primary_key"`
	UserID            int       `gorm:"null"`
	User              User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	AccountName       string    `gorm:"type:varchar(255)" validate:"required,min=3,max=255"`
	AccountCode       string    `gorm:"type:varchar(50)" validate:"required,min=3"`
	AccountNumber     int64     `gorm:"type:bigint;unique" validate:"required"`
	AccountHolderName string    `gorm:"type:varchar(255)" validate:"required,min=3,max=255"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}
