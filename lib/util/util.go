package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3
	"github.com/pkg/errors"
)

// DbOpen は データベースを開きます
func DbOpen() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		errors.Wrap(err, "dbに接続できませんでした")
	}
	db.LogMode(true)
	return db, nil
}
