package model

import (
	"github.com/jinzhu/gorm"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

// UserModel is model struct of user
type UserModel struct {
	db *gorm.DB
}

// NewUserModel creates UserModel
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db}
}

// IndexResponse is response format for Index
type IndexResponse struct {
	Users []*schema.User `json:"users"`
}

// Index returns all users
func (u *UserModel) Index() IndexResponse {
	users := []*schema.User{}
	u.db.Find(&users)
	return IndexResponse{Users: users}
}
