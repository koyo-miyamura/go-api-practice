package model

import (
	"github.com/jinzhu/gorm"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

// UserModel is interface of UserModel
type UserModel interface {
	Index() *IndexResponse
	Show(id uint64) *ShowResponse
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

// ShowResponse is response format for Show
type ShowResponse struct {
	User *schema.User `json:"user"`
}

// Show returns all users
func (u *userModel) Show(id uint64) *ShowResponse {
	user := &schema.User{}
	u.db.Find(&user, id)
	return &ShowResponse{User: user}
}
