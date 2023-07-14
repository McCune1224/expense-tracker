package store

import (
	"github.com/mccune1224/expense-tracker/model"
	"gorm.io/gorm"
)

// Really only using interface here incase I change my mind on DB later
type UserStore interface {
	// Find user by some query (email, id, username, etc)
	GetUser(field, value string) (*model.User, error)
	// Find all users by some query (email, id, username, etc)
	GetUsers(query string) (*[]model.User, error)
	// Create a new user in the database
	CreateUser(user *model.User) error
	// Update a user in the database
	UpdateUser(user *model.User) error
	// Delete a user from the database
	DeleteUser(user *model.User) error
}

type PostgreUserStore struct {
	db *gorm.DB
}

func NewPostgreUserStore(db *gorm.DB) *PostgreUserStore {
	return &PostgreUserStore{
		db: db,
	}
}

// CreateUser implements UserStore.
func (pus *PostgreUserStore) CreateUser(user *model.User) error {
	err := pus.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser implements UserStore.
func (pus *PostgreUserStore) DeleteUser(user *model.User) error {
	err := pus.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (pus *PostgreUserStore) GetUser(field, value string) (*model.User, error) {
	dbUser := &model.User{}
	err := pus.db.Where(field+" = ?", value).First(&dbUser).Error
	if err != nil {
		return dbUser, err
	}
	return dbUser, nil
}

// GetUsers implements UserStore.
func (pus *PostgreUserStore) GetUsers(query string) (*[]model.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserStore.
func (pus *PostgreUserStore) UpdateUser(user *model.User) error {
	err := pus.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
