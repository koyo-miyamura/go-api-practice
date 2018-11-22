package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/koyo-miyamura/go-api-practice/model"
)

// UserHandler はどのDBを使うかを保持します
// 主にテスト用にDBを切り替えるために使用します
type UserHandler struct {
	db *gorm.DB
}

// NewUserHandler はUserHandlerを生成して返します
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

// NewUserServer create user model's handler
func (h *UserHandler) NewUserServer() *http.ServeMux {
	server := http.NewServeMux()
	server.HandleFunc("/users", h.Index)
	return server
}

// Index is user model's index
func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("/users handled")

	w.Header().Set("Content-Type", "application/json")

	um := model.NewUserModel(h.db)
	res := um.Index()

	result, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = w.Write(result)
	if err != nil {
		log.Println(err.Error())
	}
}
