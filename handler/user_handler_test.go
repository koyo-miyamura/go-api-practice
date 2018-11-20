package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
)

func TestIndex(t *testing.T) {
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	users := []model.User{
		{
			Name: "hoge",
		},
		{
			Name: "fuga",
		},
	}
	for _, user := range users {
		db.Create(&user)
	}

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	h := NewUserHandler(db)
	h.Index(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status code %v", w.Code)
	}

	res := &IndexResponse{}
	util.JSONRead(w, res)

	want := users
	got := res.Users

	if len(got) != len(want) {
		t.Fatalf("error length got %v, want %v", len(got), len(want))
	}
	for i := 0; i < len(want); i++ {
		if want[i].Name != got[i].Name {
			t.Errorf("got user name %v, want %v", got[i].Name, want[i].Name)
		}
	}
}
