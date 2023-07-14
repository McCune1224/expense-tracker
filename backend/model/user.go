package model

import "gorm.io/gorm"

// Gorm Database Representation of a User
type User struct {
	gorm.Model
	Username     string
	Password     string
	Balance      int
	Email        string
	Transactions []Transaction
	// Default categories provided, but user can add more
	Categories []Category
}
