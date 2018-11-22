package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/koyo-miyamura/go-api-practice/lib/util"
	"github.com/koyo-miyamura/go-api-practice/model"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

func TestIndex(t *testing.T) {
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	users := []schema.User{
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

	um := model.NewUserModel(db)
	h := NewUserHandler(um)
	h.Index(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status code %v", w.Code)
	}

	res := &model.IndexResponse{}
	util.JSONRead(w, res)

	want := users
	got := res.Users

	if len(got) != len(want) {
		t.Fatalf("length got %v, want %v", len(got), len(want))
	}
	for i := 0; i < len(want); i++ {
		if want[i].Name != got[i].Name {
			t.Errorf("user name got %v, want %v", got[i].Name, want[i].Name)
		}
	}
}
