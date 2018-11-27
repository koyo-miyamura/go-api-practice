package stub

import (
	"errors"

	"github.com/koyo-miyamura/go-api-practice/model"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

// UserModel is xxx
type UserModel struct {
	IndexStub    func() *model.IndexResponse
	ShowStub     func(id uint64) (*model.ShowResponse, error)
	CreateStub   func(user *schema.User) (*model.CreateResponse, error)
	UpdateStub   func(user *schema.User) (*model.UpdateResponse, error)
	DeleteStub   func(id uint64) error
	ValidateStub func(user *schema.User) error
}

// Index return stub of UserModel.Index
func (u *UserModel) Index() *model.IndexResponse {
	return u.IndexStub()
}

// Show return stub of UserModel.Show
func (u *UserModel) Show(id uint64) (*model.ShowResponse, error) {
	return u.ShowStub(id)
}

// Create return stub of UserModel.Create
func (u *UserModel) Create(user *schema.User) (*model.CreateResponse, error) {
	return u.CreateStub(user)
}

// Update return stub of UserModel.Update
func (u *UserModel) Update(user *schema.User) (*model.UpdateResponse, error) {
	return u.UpdateStub(user)
}

// Delete return stub of UserModel.Delete
func (u *UserModel) Delete(id uint64) error {
	return u.DeleteStub(id)
}

// Validate is stub of UserModel.Validate
func (u *UserModel) Validate(user *schema.User) error {
	return u.ValidateStub(user)
}

// NewUserModel returns stub of UserModel
func NewUserModel() *UserModel {
	return &UserModel{
		IndexStub: func() *model.IndexResponse {
			return nil
		},
		ShowStub: func(id uint64) (*model.ShowResponse, error) {
			return nil, errors.New("not implementation")
		},
		CreateStub: func(user *schema.User) (*model.CreateResponse, error) {
			return nil, errors.New("not implementation")
		},
		UpdateStub: func(user *schema.User) (*model.UpdateResponse, error) {
			return nil, errors.New("not implementation")
		},
		DeleteStub: func(id uint64) error {
			return nil
		},
		ValidateStub: func(user *schema.User) error {
			return nil
		},
	}
}
