package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm用
	"github.com/koyo-miyamura/go-api-practice/model"
)

// Responce はJsonレスポンスのフォーマット
type Responce struct {
	Users []model.User `json:"users"`
}

// NewUserServer create user model's handler
func NewUserServer() *http.ServeMux {
	server := http.NewServeMux()
	server.HandleFunc("/", UserIndex)
	return server
}

// UserIndex is user model's index
func UserIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("/ handled")
	w.Header().Set("Content-Type", "application/json")

	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		panic("dbに接続できませんでした")
	}
	defer db.Close()
	db.LogMode(true)

	users := []model.User{}
	db.Find(&users)
	res := Responce{
		Users: users,
	}

	result, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = w.Write(result)
	if err != nil {
		log.Println(err.Error())
	}
}
