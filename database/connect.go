package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect establishes a connection to the SQLite database.
func Connect() (DB *gorm.DB, Error error) {
	if db == nil {
		var err error
		// Open a connection to the SQLite database
		db, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("Failed to open DB")
		}
	}

	return db, nil
}
