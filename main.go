package main

import (
	"log"
	"net/http"

	"github.com/koyo-miyamura/go-api-practice/handler"
)

func main() {
	log.Println("Server started!")
	userServer := handler.NewUserServer()
	log.Fatal(http.ListenAndServe(":8080", userServer))
}
