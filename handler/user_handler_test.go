package handler

import (
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
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

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
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()

	want := &model.ShowResponse{
		User: &schema.User{
			ID:   1,
			Name: "hoge",
		},
	}

	um := stub.NewUserModel()
	um.ShowStub = func(id uint64) *model.ShowResponse {
		if id == 1 {
			return want
		}
		return nil
	}
	h := NewUserHandler(um)
	r := h.NewUserServer()
	r.ServeHTTP(w, req) // gorilla/muxがパラメータ取り出すにはこれを経由させる必要がある

	if w.Code != http.StatusOK {
		t.Fatalf("status code %v", w.Code)
	}

	got := &model.ShowResponse{}
	util.JSONRead(w, got)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("responce got %v, want %v", got, want)
	}

	// // Fail
	// req = httptest.NewRequest(http.MethodGet, "/users/10000", nil)
	// h.Show(w, req)

	// if w.Code != http.StatusNotFound {
	// 	t.Fatalf("status code %v", w.Code)
	// }
}
