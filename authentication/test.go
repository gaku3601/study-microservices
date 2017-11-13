package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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

type JwtKey struct {
	Created_at  int    `json:"created_at"`
	Id          string `json:"id"`
	Algorithm   string `json:"algorithm"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	Consumer_id string `json:"consumer_id"`
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
		//jwtの問い合わせ
		url := "http://localhost:8001/consumers/gaku/jwt"
		req, _ := http.NewRequest(
			"POST",
			url,
			nil,
		)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		jwtKey := new(JwtKey)
		client := &http.Client{}
		resp, _ := client.Do(req)
		json.NewDecoder(resp.Body).Decode(&jwtKey)
		defer resp.Body.Close()

		createTokenString(jwtKey)

	} else {
		user.IsAuth = false
	}
	json.NewEncoder(w).Encode(user)
}

func createTokenString(jwtKey *JwtKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": jwtKey.Key,
	})
	tokenString, _ := token.SignedString([]byte(jwtKey.Secret))

	fmt.Println(tokenString)
	return "a"
}
