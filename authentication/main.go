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
	r.HandleFunc("/users/login", Login).Methods("POST")

	http.ListenAndServe(":8080", r)
}

type User struct {
	ID   string `json:"id"`
	Pass string `json:"pass"`
}

type Response struct {
	Token string `json:"token"`
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillaaaaaaaaadaaaaaaaa!\n"))
}

//ログイン認証。ログイン完了後、JWTトークンを返却する
func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	error := decoder.Decode(&user)
	if error != nil {
		w.Write([]byte("json decode error " + error.Error() + "\n"))
	}

	//認証処理
	res := new(Response)
	if user.ID == "gaku" && user.Pass == "gakugaku" {
		//jwtの問い合わせ
		res.Token = fetchCreateToken()

	} else {
		res.Token = ""
	}
	//返却
	json.NewEncoder(w).Encode(res)
}

//ユーザ登録
func SignUp(w http.ResponseWriter, r *http.Request) {
}
