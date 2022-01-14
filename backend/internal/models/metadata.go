package models

import "time"

type Metadata struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Key       string    `json:"key" gorm:"unique"`
	Value     string    `json:"value"`
}

var KEY_SITE_NAME = "site name"
var KEY_OWNER_PASSWORD = "owner password"
