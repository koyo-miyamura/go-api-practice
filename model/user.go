package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/koyo-miyamura/go-api-practice/schema"
)

// UserModel is interface of UserModel
type UserModel interface {
	Index() *IndexResponse
	Show(id uint64) (*ShowResponse, error)
	Create(req *CreateRequest) (*CreateResponse, error)
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
func (u *userModel) Show(id uint64) (*ShowResponse, error) {
	user := &schema.User{}
	if err := u.db.Find(&user, id).Error; err != nil {
		return nil, errors.New("error find user")
	}
	return &ShowResponse{User: user}, nil
}

// CreateRequest is request format for Create
type CreateRequest struct {
	User *schema.User `json:"user"`
}

// CreateResponse is response format for Create
type CreateResponse struct {
	User *schema.User `json:"user"`
}

// Create creates new user
func (u *userModel) Create(req *CreateRequest) (*CreateResponse, error) {
	user := req.User
	if err := u.db.Create(&user).Error; err != nil {
		return nil, errors.New("error create user")
	}
	res := &CreateResponse{
		User: user,
	}
	return res, nil
}
