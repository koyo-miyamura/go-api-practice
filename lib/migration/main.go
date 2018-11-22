package main

import (
	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

func main() {
	db, err := util.DbOpen()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&schema.User{})
	db.Model(&schema.User{}).AddIndex("idx_user_name", "name")
}
