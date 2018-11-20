package model

// Models は全モデルのスライスを返します
func Models() []interface{} {
	return []interface{}{
		&User{},
	}
}
