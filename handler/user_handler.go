package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/koyo-miyamura/go-api-practice/model"
)

// Responce はJsonレスポンスのフォーマット
type Responce struct {
	Status uint32       `json:"status"`
	Users  []model.User `json:"users"`
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
	res := Responce{
		Status: http.StatusOK,
		Users: []model.User{
			{
				UserID: 1,
				Name:   "hoge",
			}, {
				UserID: 2,
				Name:   "fuga",
			},
		},
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
