package stub

import (
	"github.com/koyo-miyamura/go-api-practice/model"
)

// UserModel is xxx
type UserModel struct {
	IndexStub func() *model.IndexResponse
}

// Index return stub of UserModel
func (u *UserModel) Index() *model.IndexResponse {
	return u.IndexStub()
}

// NewUserModel returns stub of UserModel
func NewUserModel() *UserModel {
	return &UserModel{
		IndexStub: func() *model.IndexResponse {
			return nil
		},
	}
}
