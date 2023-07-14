package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/mccune1224/expense-tracker/store"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
	us store.UserStore
	ts store.TransactionStore
	v  *validator.Validate
}

func New(gormDB *gorm.DB, us store.UserStore, ts store.TransactionStore, migrations ...interface{}) *Handler {
	if len(migrations) > 0 {
		err := gormDB.AutoMigrate(migrations...)
		if err != nil {
			log.Println("Failed to migrate model, error: " + err.Error())
		}
	}
	return &Handler{
		db: gormDB,
		us: us,
		ts: ts,
		v:  validator.New(),
	}
}
