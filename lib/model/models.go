package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"PRIMARY_KEY;Type:varchar(25);UNIQUE;INDEX" validate:"max=25,ascii"`
	Password string `gorm:"Type:varchar(100);NOT NULL"`
	Name     string `gorm:"Type:varchar(20);NOT NULL"`
}
