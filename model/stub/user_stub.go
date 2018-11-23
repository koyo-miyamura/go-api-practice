package stub

import (
	"github.com/koyo-miyamura/go-api-practice/model"
)

// UserModel is xxx
type UserModel struct {
	IndexStub func() *model.IndexResponse
	ShowStub  func(id uint64) *model.ShowResponse
}

// Index return stub of UserModel.Index
func (u *UserModel) Index() *model.IndexResponse {
	return u.IndexStub()
}

// Show return stub of UserModel.Show
func (u *UserModel) Show(id uint64) *model.ShowResponse {
	return u.ShowStub(id)
}

// NewUserModel returns stub of UserModel
func NewUserModel() *UserModel {
	return &UserModel{
		IndexStub: func() *model.IndexResponse {
			return nil
		},
		ShowStub: func(id uint64) *model.ShowResponse {
			return nil
		},
	}
}
