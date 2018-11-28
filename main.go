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

	// CORS対策
	// 参考：https://github.com/gorilla/handlers/blob/master/cors.go#L30
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders(
		[]string{"Accept", "Accept-Language", "Content-Language", "Content-type", "Origin"},
	)
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMiddleWare := handlers.CORS(allowedMethods, allowedHeaders, allowedOrigins)
	server := corsMiddleWare(userServer)

	log.Fatal(http.ListenAndServe(":8080", server))
}
