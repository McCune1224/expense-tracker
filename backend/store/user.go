package store

import (
	"github.com/mccune1224/expense-tracker/model"
	"gorm.io/gorm"
)

// Really only using interface here incase I change my mind on DB later
type UserStore interface {
	// Find user by some query (email, id, username, etc)
	GetUser(query string)
	// Find all users by some query (email, id, username, etc)
	GetUsers(query string)
	// Create a new user in the database
	CreateUser(user model.User)
	// Update a user in the database
	UpdateUser(user model.User)
	// Delete a user from the database
	DeleteUser(user model.User)
}

type PostgreUserStore struct {
	db *gorm.DB
}

func (pus *PostgreUserStore) GetUser()    {}
func (pus *PostgreUserStore) GetUsers()   {}
func (pus *PostgreUserStore) CreateUser() {}
func (pus *PostgreUserStore) UpdateUser() {}
func (pus *PostgreUserStore) DeleteUser() {}
