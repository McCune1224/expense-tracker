package store

import "gorm.io/gorm"

type PostgreUserStore struct {
	db *gorm.DB
}

func (pus *PostgreUserStore) GetUser()    {}
func (pus *PostgreUserStore) GetUsers()   {}
func (pus *PostgreUserStore) CreateUser() {}
func (pus *PostgreUserStore) UpdateUser() {}
func (pus *PostgreUserStore) DeleteUser() {}
