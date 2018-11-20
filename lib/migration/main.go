package main

import (
	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
)

func main() {
	db, err := util.DbOpen()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&model.User{})
	db.AutoMigrate(&model.User{})
	db.Model(&model.User{}).AddIndex("idx_user_name", "name")
}