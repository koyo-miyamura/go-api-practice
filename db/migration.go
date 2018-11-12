package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/koyo-miyamura/go-api-practice/model"
)

func main() {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		panic("dbに接続できませんでした")
	}
	defer db.Close()
	db.LogMode(true)

	// マイグレーション
	db.DropTableIfExists(&model.User{})
	db.AutoMigrate(&model.User{})
	db.Model(&model.User{}).AddIndex("idx_user_name", "name")

	// 初期seedも設定(実行ごとに挿入)
	users := []model.User{
		{
			Name: "hoge",
		}, {
			Name: "fuga",
		},
	}
	for _, user := range users {
		if db.NewRecord(user) {
			db.Create(&user)
		}
	}
}
