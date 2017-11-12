package main

import (
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux.git"
)

func main() {
	r := mux.NewRouter()
	// 単純なハンドラ
	r.HandleFunc("/", YourHandler).Methods("POST")
	r.HandleFunc("/users/auth", UserAuth).Methods("POST")

	http.ListenAndServe(":8080", r)
}

type User struct {
	ID     string `json:"id"`
	Pass   string `json:"pass"`
	IsAuth bool   `json:"isauth"`
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillaaaaaaaaadaaaaaaaa!\n"))
}

func UserAuth(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	error := decoder.Decode(&user)
	if error != nil {
		w.Write([]byte("json decode error " + error.Error() + "\n"))
	}

	//認証処理
	if user.ID == "gaku" && user.Pass == "gakugaku" {
		user.IsAuth = true
	} else {
		user.IsAuth = false
	}
	json.NewEncoder(w).Encode(user)
}
