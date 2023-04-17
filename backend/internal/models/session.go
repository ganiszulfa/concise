package models

import "time"

type Session struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	ExpiredAt time.Time `json:"expiredAt"`
	IsDeleted bool      `json:"isDeleted" gorm:"default:false"`
}
