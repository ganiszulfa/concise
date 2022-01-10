package models

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Email     string    `json:"name" gorm:"unique"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	IsOwner   bool      `json:"isOwner" gorm:"default:false"`

	Token string `json:"token,omitempty" gorm:"-"`
}
