package database

import (
	"github.com/jinzhu/gorm"
	"smooth-bear.live/lib/model"
)

type Accessor interface {
	CreateUser(user *model.User) (resultUser *model.User, err error)

	// For Transaction
	BeginTx()
	Commit() *gorm.DB
	Rollback() *gorm.DB
}
