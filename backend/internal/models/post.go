package models

import "time"

type Post struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Title       string    `json:"title" gorm:"type:varchar(200)"`
	Slug        string    `json:"slug" gorm:"uniqueIndex;type:varchar(200)"`
	Content     string    `json:"value" gorm:"type:text"`
	PublishedAt time.Time `json:"publishedAt" gorm:"default:null"`
	IsPublished bool      `json:"isPublished" gorm:"default:false"`
	IsPage      bool      `json:"isPage" gorm:"default:false"`
	IsDeleted   bool      `json:"isDeleted" gorm:"default:false"`
}
