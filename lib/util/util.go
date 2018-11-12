package util

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3
)

// DbOpen は データベースを開きます
func DbOpen() *gorm.DB {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		panic("dbに接続できませんでした")
	}
	db.LogMode(true)
	return db
}
