package models

import (
	"gorm.io/gorm"
)

func AutoMigrateAllTables(db *gorm.DB) {
	err := db.AutoMigrate(User{}, Metadata{})
	if err != nil {
		panic(err)
	}
}
