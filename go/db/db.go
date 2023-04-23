package db

import (
	"mercari-build-training-2022/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitialiseDB() *gorm.DB {
	// Initialise DB
	db, err := gorm.Open(sqlite.Open("../../db/items.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Category{})

	return db
}
