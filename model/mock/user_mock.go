package mock

import (
	"github.com/koyo-miyamura/go-api-practice/model"
)

// UserModel is xxx
type UserModel struct {
	IndexMock func() *model.IndexResponse
}

// Index return mock of UserModel
func (u *UserModel) Index() *model.IndexResponse {
	return u.IndexMock()
}

// NewUserModel returns mock of UserModel
func NewUserModel() *UserModel {
	return &UserModel{
		IndexMock: func() *model.IndexResponse {
			return nil
		},
	}
}
