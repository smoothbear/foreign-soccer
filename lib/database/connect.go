package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectToMysql() (db *gorm.DB, err error) {
	args := os.Getenv("DB_URL")

	db, err = gorm.Open("mysql", args)
	return
}
