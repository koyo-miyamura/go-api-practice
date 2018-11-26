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

	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content-type got %#v, want %#v", contentType, "application/json")
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

		if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
			t.Errorf("Content-type got %#v, want %#v", contentType, "application/json")
		}

		got := &model.ShowResponse{}
		if err := util.JSONRead(w, got); err != nil {
			t.Fatal(err)
		}

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
	user := &schema.User{
		Name: "hoge",
	}
	request := &CreateRequest{
		Name: user.Name,
	}
	want := &model.CreateResponse{
		User: user,
	}

	type Test struct {
		Title      string
		Create     bool
		Validate   bool
		StatusCode int
	}
	tests := []Test{
		{
			Title:      "Success",
			Create:     true,
			Validate:   true,
			StatusCode: http.StatusCreated,
		},
		{
			Title:      "Validate false",
			Create:     true,
			Validate:   false,
			StatusCode: http.StatusBadRequest,
		},
		{
			Title:      "Create false",
			Create:     false,
			Validate:   true,
			StatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			um := stub.NewUserModel()
			um.CreateStub = func(user *schema.User) (*model.CreateResponse, error) {
				if test.Create {
					return want, nil
				}
				return nil, errors.New("create error")
			}
			um.ValidateStub = func(user *schema.User) error {
				if test.Validate {
					return nil
				}
				return errors.New("validate error")
			}

			h := NewUserHandler(um)
			r := h.NewUserServer()

			input, err := json.Marshal(request)
			if err != nil {
				t.Fatal("error Marshal test input")
			}
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(input))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != test.StatusCode {
				t.Errorf("status code got %v, want %v", w.Code, test.StatusCode)
			}

			if test.Create && test.Validate {
				got := &model.CreateResponse{}
				util.JSONRead(w, got)

				if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
					t.Errorf("Content-type got %#v, want %#v", contentType, "application/json")
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("responce got %v, want %v", got, want)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	um := stub.NewUserModel()
	um.DeleteStub = func(id uint64) error {
		if id == 1 {
			return nil
		}
		return errors.New("Delete stub error")
	}

	h := NewUserHandler(um)
	r := h.NewUserServer()

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if want := http.StatusNoContent; w.Code != want {
			t.Fatalf("status code %v, want %v", w.Code, want)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/2", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if want := http.StatusInternalServerError; w.Code != want {
			t.Fatalf("status code %v, want %v", w.Code, want)
		}
	})
}
