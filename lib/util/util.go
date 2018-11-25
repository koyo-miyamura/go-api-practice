package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite3
	"github.com/koyo-miyamura/go-api-practice/schema"
	"github.com/pkg/errors"
)

var (
	gopath = os.Getenv("GOPATH")
	home   = filepath.Join(gopath, "src", "github.com", "koyo-miyamura", "go-api-practice")
)

// DbOpen は データベースを開きます
func DbOpen() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", home+"/sqlite3/database.db")
	if err != nil {
		return nil, errors.Wrap(err, "dbに接続できませんでした")
	}
	db.LogMode(true)
	return db, nil
}

// TestDbNew はテスト用のDBを準備します
func TestDbNew() (*gorm.DB, error) {

	db, err := gorm.Open("sqlite3", home+"/sqlite3/database_test.db")
	if err != nil {
		return nil, errors.Wrap(err, "dbに接続できませんでした")
	}
	db.LogMode(true)

	// テスト用のdbをmigrate
	for _, model := range schema.Models() {
		db.AutoMigrate(model)
	}

	return db, nil
}

// TestDbClose はテスト用のDBを消去してCloseします
func TestDbClose(db *gorm.DB) {
	for _, model := range schema.Models() {
		db.Delete(model)
	}
	db.Close()
}

// JSONRead はhttptest.ResponseRecorderからJsonを読み取ります
// テスト用の関数
func JSONRead(w *httptest.ResponseRecorder, res interface{}) error {
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadAllに失敗しました")
	}

	if err := json.Unmarshal(body, res); err != nil {
		return errors.Wrap(err, "Unmarshalに失敗しました")
	}

	return nil
}

// JSONWrite は与えられた形式でJsonレスポンスを返します
func JSONWrite(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}

// ScanRequest はJsonリクエストからrequestの方に合わせてUnmarshalします
func ScanRequest(r *http.Request, request interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, request); err != nil {
		return err
	}

	return nil
}
