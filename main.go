package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User のレスポンスフォーマット
type User struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

// Responce はJsonレスポンスのフォーマット
type Responce struct {
	Status uint32 `json:"status"`
	Users  []User `json:"users"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("/ handled")
	w.Header().Set("Content-Type", "application/json")
	res := Responce{
		Status: http.StatusOK,
		Users: []User{
			{
				UserID: 1,
				Name:   "hoge",
			}, {
				UserID: 2,
				Name:   "fuga",
			},
		},
	}
	result, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = w.Write(result)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	log.Println("Server started!")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
