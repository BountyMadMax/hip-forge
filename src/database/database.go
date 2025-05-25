package database

import (
	"hip-forge/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/hip-forge.db"), &gorm.Config{})

	if err != nil {
		panic("Error opening db")
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Account{},
		&models.DNSRecord{},
	)
}
