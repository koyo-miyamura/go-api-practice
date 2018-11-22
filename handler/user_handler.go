package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/koyo-miyamura/go-api-practice/model"
)

// UserHandler はどのmodelを使うかを保持します
// 主にテスト用にmodelを切り替えるために使用します
type UserHandler struct {
	model model.UserModel
}

// NewUserHandler はUserHandlerを生成して返します
func NewUserHandler(m model.UserModel) *UserHandler {
	return &UserHandler{m}
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

	res := h.model.Index()

	result, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = w.Write(result)
	if err != nil {
		log.Println(err.Error())
	}
}
