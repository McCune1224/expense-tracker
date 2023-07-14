package model

import (
	"time"

	"gorm.io/gorm"
)

// Gorm Database Representation of a Transaction
type Transaction struct {
	gorm.Model
	Name   string
	Amount int
	Date   time.Time
	Note   string
	// FK to User (has many)
	UserID   uint
	Category Category
}

type Category struct {
	gorm.Model
	Name string
	Icon string // Emoji or Font Awesome Icon?
	// FK to Transaction (has one)
	TransactionID uint
	// FK to User (has many)
	UserID uint
}
