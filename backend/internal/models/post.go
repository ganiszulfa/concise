package models

import "time"

type Post struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug" gorm:"unique"`
	Content     string    `json:"value" gorm:"type:text"`
	AuthorID    int       `json:"authorId"`
	Author      User      `json:"author"`
	IsPublished bool      `json:"isPublished" gorm:"default:false"`
	PublishedAt time.Time `json:"publishedAt" gorm:"default:null"`
}
