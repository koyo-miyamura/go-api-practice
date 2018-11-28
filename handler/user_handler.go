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
	router.HandleFunc("/users/{id:[0-9]+}", h.Update).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", h.Delete).Methods("DELETE")

	// Handle VueAPP
	router.PathPrefix("/_nuxt").Handler(http.FileServer(http.Dir("public/dist")))
	router.PathPrefix("/").HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/dist/index.html")
	}))
	return router
}

// Index is user model's index
func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("/users handled")

	res := h.model.Index()

	if err := util.JSONWrite(w, res, http.StatusOK); err != nil {
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
		log.Println(errors.Wrapf(err, "error parse uint:%v", idStr))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("/users/%d handled", id)

	res, err := h.model.Show(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := util.JSONWrite(w, res, http.StatusOK); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// CreateRequest is request format for Create
type CreateRequest struct {
	Name string `json:"name"`
}

// Create is user model's create
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Print("/users POST handled")

	req := &CreateRequest{}
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

	if err := util.JSONWrite(w, res, http.StatusCreated); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// UpdateRequest is request format for Create
type UpdateRequest struct {
	Name string `json:"name"`
}

// Update is user model's update
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println(errors.Wrapf(err, "error parse uint:%v", idStr))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("/users/%d PUT handled", id)

	req := &UpdateRequest{}
	if err := util.ScanRequest(r, req); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &schema.User{
		ID:   id,
		Name: req.Name,
	}

	if err := h.model.Validate(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.model.Update(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := util.JSONWrite(w, res, http.StatusOK); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Delete is user model's delete
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println(errors.Wrapf(err, "error parse uint:%v", idStr))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("/users/%d DELETE handled", id)

	if err := h.model.Delete(id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
