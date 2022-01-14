package models

import (
	"gorm.io/gorm"
)

func AutoMigrateAllTables(db *gorm.DB) {
	err := db.AutoMigrate(
		User{},
		Metadata{},
		Post{},
	)
	if err != nil {
		panic(err)
	}
}
