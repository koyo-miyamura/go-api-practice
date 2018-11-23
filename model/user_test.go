package model

import (
	"fmt"
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

	users := []schema.User{}
	for i := 0; i < 10; i++ {
		users = append(users, schema.User{Name: fmt.Sprintf("hoge%d", i)})
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

func TestShow(t *testing.T) {
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	user := schema.User{
		ID:   1,
		Name: "hoge",
	}
	db.Create(&user)

	um := NewUserModel(db)
	res := um.Show(1)

	got := res.User
	want := user

	if want.ID != got.ID {
		t.Errorf("user ID got %v, want %v", got.ID, want.ID)
	}

	if want.Name != got.Name {
		t.Errorf("user name got %v, want %v", got.Name, want.Name)
	}
}
