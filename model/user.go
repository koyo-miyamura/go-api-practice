package model

import "github.com/jinzhu/gorm"

// User モデルの定義
type User struct {
	gorm.Model
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}
