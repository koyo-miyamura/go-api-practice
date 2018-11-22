package handler

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
	"github.com/koyo-miyamura/go-api-practice/model/mock"
)

func TestIndex(t *testing.T) {
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	um := mock.NewUserModel()
	want := &model.IndexResponse{}
	um.IndexMock = func() *model.IndexResponse {
		return want
	}
	h := NewUserHandler(um)
	h.Index(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status code %v", w.Code)
	}

	got := &model.IndexResponse{}
	util.JSONRead(w, got)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("responce got %v, want %v", got, want)
	}
}
