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
	user := &schema.User{
		Name: "hoge",
	}
	req := &model.CreateRequest{
		Name: user.Name,
	}
	input, err := json.Marshal(req)
	if err != nil {
		t.Fatal("error Marshal test input")
	}
	want := &model.CreateResponse{
		User: user,
	}

	type Test struct {
		Title      string
		Input      []byte
		Create     bool
		Validate   bool
		StatusCode int
	}
	tests := []Test{
		{
			Title:      "Success",
			Input:      input,
			Create:     true,
			Validate:   true,
			StatusCode: http.StatusOK,
		},
		{
			Title:      "Validate false",
			Input:      input,
			Create:     true,
			Validate:   false,
			StatusCode: http.StatusBadRequest,
		},
		{
			Title:      "Create false",
			Input:      input,
			Create:     false,
			Validate:   true,
			StatusCode: http.StatusInternalServerError,
		},
		{
			Title:      "nil input",
			Input:      nil,
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

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(test.Input))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != test.StatusCode {
				t.Errorf("status code got %v, want %v", w.Code, test.StatusCode)
			}

			if test.Create && test.Validate {
				got := &model.CreateResponse{}
				util.JSONRead(w, got)

				if !reflect.DeepEqual(got, want) {
					t.Errorf("responce got %v, want %v", got, want)
				}
			}
		})
	}
}
