package model

import "time"

// User モデルの定義
type User struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
