package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
	"github.com/koyo-miyamura/go-api-practice/model/stub"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

func TestIndex(t *testing.T) {
	want := &model.IndexResponse{
		Users: []*schema.User{
			{
				Name: "hoge",
			},
			{
				Name: "fuga",
			},
		},
	}

	um := stub.NewUserModel()
	um.IndexStub = func() *model.IndexResponse {
		return want
	}

	h := NewUserHandler(um)
	r := h.NewUserServer()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status code %v", w.Code)
	}

	got := &model.IndexResponse{}
	util.JSONRead(w, got)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("responce got %v, want %v", got, want)
	}
}

func TestShow(t *testing.T) {
	want := &model.ShowResponse{
		User: &schema.User{
			ID:   1,
			Name: "hoge",
		},
	}

	um := stub.NewUserModel()
	um.ShowStub = func(id uint64) (*model.ShowResponse, error) {
		if id == 1 {
			return want, nil
		}
		return nil, errors.New("can't find user")
	}

	h := NewUserHandler(um)
	r := h.NewUserServer()

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req) // gorilla/muxがパラメータ取り出すにはこれを経由させる必要がある

		if w.Code != http.StatusOK {
			t.Fatalf("status code %v", w.Code)
		}

		got := &model.ShowResponse{}
		util.JSONRead(w, got)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("responce got %v, want %v", got, want)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("status code %v", w.Code)
		}
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		validUser := &schema.User{
			Name: "hoge",
		}
		input := &model.CreateRequest{
			Name: validUser.Name,
		}
		want := &model.CreateResponse{
			User: validUser,
		}
		um := stub.NewUserModel()
		um.CreateStub = func(req *model.CreateRequest) (*model.CreateResponse, error) {
			res := &model.CreateResponse{
				User: validUser,
			}
			return res, nil
		}

		h := NewUserHandler(um)
		r := h.NewUserServer()

		jsonInput, err := json.Marshal(input)
		if err != nil {
			t.Fatal("error Marshal test input")
		}
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonInput))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("status code %v", w.Code)
		}

		got := &model.CreateResponse{}
		util.JSONRead(w, got)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("responce got %v, want %v", got, want)
		}
	})
}
