package model

import (
	"github.com/jinzhu/gorm"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

// UserModel is interface of UserModel
type UserModel interface {
	Index() *IndexResponse
}

// userModel is model struct of user
type userModel struct {
	db *gorm.DB
}

// NewUserModel creates UserModel
func NewUserModel(db *gorm.DB) UserModel {
	return &userModel{db}
}

// IndexResponse is response format for Index
type IndexResponse struct {
	Users []*schema.User `json:"users"`
}

// Index returns all users
func (u *userModel) Index() *IndexResponse {
	users := []*schema.User{}
	u.db.Find(&users)
	return &IndexResponse{Users: users}
}
