package handler

import (
	"log"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
	// UserStore *model.UserStore
}

func New(gormDB *gorm.DB, migrations ...interface{}) *Handler {
	if len(migrations) > 0 {
		err := gormDB.AutoMigrate(migrations...)
		if err != nil {
			log.Println("Failed to migrate model, error: " + err.Error())
		}
	}
	return &Handler{
		db: gormDB,
	}
}

