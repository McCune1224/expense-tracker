package model

import "gorm.io/gorm"

// Database representation of a User
type User struct {
	gorm.Model
	Username     string `gorm:"unique,not null"`
	Password     string `gorm:"not null,"`
	Balance      int
	Email        string `gorm:"unique,not null"`
	Transactions []Transaction
	// Default categories provided, but user can add more
	Categories []Category
}
