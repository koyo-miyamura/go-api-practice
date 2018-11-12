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
