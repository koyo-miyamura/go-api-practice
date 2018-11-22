package model

import (
	"testing"

	"github.com/koyo-miyamura/go-api-practice/lib/util"
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

	um := NewUserModel(db)
	res := um.Index()

	got := res.Users
	want := users

	if len(want) != len(got) {
		t.Fatalf("length got %v, want %v", len(got), len(want))
	}

	for i := 0; i < len(want); i++ {
		if want[i].Name != got[i].Name {
			t.Errorf("user name got %v, want %v", got[i].Name, want[i].Name)
		}
	}
}
