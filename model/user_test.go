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

	t.Run("Success", func(t *testing.T) {
		res, err := um.Show(1)
		if err != nil {
			t.Errorf("error Show method %v", err)
		}

		got := res.User
		want := user

		if want.ID != got.ID {
			t.Errorf("user ID got %v, want %v", got.ID, want.ID)
		}

		if want.Name != got.Name {
			t.Errorf("user name got %v, want %v", got.Name, want.Name)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		res, _ := um.Show(10000000)
		if res != nil {
			t.Errorf("want response is nil, but got %#v", res)
		}
	})
}

func TestCreate(t *testing.T) {
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	um := NewUserModel(db)

	t.Run("Success", func(t *testing.T) {
		user := &schema.User{
			Name: "hoge",
		}
		res, err := um.Create(user)
		if err != nil {
			t.Errorf("error Create method %v", err)
		}

		got := schema.User{}
		if err := db.Find(&got, res.User.ID).Error; err != nil {
			t.Fatalf("can't Find created user %v", res)
		}
		want := user

		if want.Name != got.Name {
			t.Errorf("user name got %v, want %v", got.Name, want.Name)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		res, err := um.Create(&schema.User{})
		if err != nil {
			t.Errorf("nil can't input, but got %#v", res)
		}
	})
}

func TestValidateUser(t *testing.T) {
	db, err := util.TestDbNew()
	if err != nil {
		t.Fatal(err, "DB接続できませんでした")
	}
	defer util.TestDbClose(db)

	existUser := schema.User{
		Name: "duplicated",
	}
	db.Create(&existUser)
	um := NewUserModel(db)

	type Test struct {
		Title string
		User  *schema.User
		Want  bool
	}
	tests := []Test{
		{
			Title: "Valid User",
			User: &schema.User{
				Name: "hoge",
			},
			Want: true,
		},
		{
			Title: "User having no contents",
			User:  &schema.User{},
			Want:  false,
		},
		{
			Title: "Blank Name",
			User: &schema.User{
				Name: "",
			},
			Want: false,
		},
		{
			Title: "Duplicated name",
			User: &schema.User{
				Name: existUser.Name,
			},
			Want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if err := um.Validate(test.User); (err != nil) == test.Want {
				t.Fatalf("%s must be %v", test.Title, test.Want)
			}
		})
	}
}
