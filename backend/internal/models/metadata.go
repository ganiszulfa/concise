package models

type Metadata struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Key   string `json:"key" gorm:"unique"`
	Value string `json:"value"`
}
