package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
	"github.com/koyo-miyamura/go-api-practice/schema"
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
	router.HandleFunc("/users", h.Create).Methods("POST")
	return router
}

// Index is user model's index
func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("/users handled")

	res := h.model.Index()

	if err := util.JSONWrite(w, res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
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

	res, err := h.model.Show(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := util.JSONWrite(w, res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Create is user model's create
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("/users POST handled")

	req := &model.CreateRequest{}
	if err := util.ScanRequest(r, req); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &schema.User{
		Name: req.Name,
	}

	if err := h.model.Validate(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.model.Create(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := util.JSONWrite(w, res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
