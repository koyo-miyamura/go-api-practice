package schema

import "time"

// Models は全モデルのスライスを返します
func Models() []interface{} {
	return []interface{}{
		&User{},
	}
}

// User モデルの定義
type User struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `gorm:"unique;not null;index:idx_user_name" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
