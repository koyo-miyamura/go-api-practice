package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/koyo-miyamura/go-api-practice/handler"
	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
)

func main() {
	log.Println("Server started!")

	db, err := util.DbOpen()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userModel := model.NewUserModel(db)
	userHandler := handler.NewUserHandler(userModel)
	userServer := userHandler.NewUserServer()

	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"})
	corsMiddleWare := handlers.CORS(allowedMethods)
	server := corsMiddleWare(userServer)

	log.Fatal(http.ListenAndServe(":8080", server))
}
