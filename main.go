package main

import (
	"log"
	"net/http"

	"github.com/koyo-miyamura/go-api-practice/handler"
	"github.com/koyo-miyamura/go-api-practice/lib/util"
)

func main() {
	log.Println("Server started!")

	db, err := util.DbOpen()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler := handler.NewUserHandler(db)
	userServer := userHandler.NewUserServer()
	log.Fatal(http.ListenAndServe(":8080", userServer))
}
