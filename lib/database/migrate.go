package database

import (
	"github.com/jinzhu/gorm"
	"smooth-bear.live/lib/model"
)

func Migrate(db *gorm.DB) {
	db.LogMode(false)

	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
}
