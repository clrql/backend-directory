package models

import (
	"time"

	"gorm.io/gorm"
)

// UserModel represents the structure of a user model.
type UserModel struct {
	gorm.Model            // Embedded GORM model
	Name        string    // User's name
	Birth       time.Time // User's birthdate
	Address     string    // User's address
	Description *string   // User's description (optional)
}
