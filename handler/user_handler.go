package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/koyo-miyamura/go-api-practice/model"
	"github.com/pkg/errors"
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
func (h *UserHandler) NewUserServer() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", h.Index).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", h.Show).Methods("GET")
	return router
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

// Show is user model's show
func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Panicln(errors.Wrapf(err, "error parse uint:%v", idStr))
	}

	log.Printf("/users/%d handled", id)

	w.Header().Set("Content-Type", "application/json")

	res, err := h.model.Show(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
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
