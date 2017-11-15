package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	mux "github.com/gorilla/mux.git"
	viper "github.com/spf13/viper.git"
)

func main() {
	SetConfig()

	r := mux.NewRouter()
	// 単純なハンドラ
	r.HandleFunc("/", YourHandler).Methods("POST")
	r.HandleFunc("/users/login", Login).Methods("POST")
	r.HandleFunc("/users/signup", SignUp).Methods("POST")

	http.ListenAndServe(":8080", r)
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	if user.Email == "gaku" && user.Password == "gakugaku" {
		//jwtトークンの取得
		res.Token = fetchCreateToken()
	} else {
		res.Token = ""
	}
	//返却
	json.NewEncoder(w).Encode(res)
}

//ユーザ登録
func SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := new(User)

	error := decoder.Decode(&user)
	if error != nil {
		w.Write([]byte("json decode error " + error.Error() + "\n"))
	}

	//DB登録
	db, _ := sql.Open("postgres", "user=postgres host=localhost dbname=auth_db port=5433 sslmode=disable")
	defer db.Close()

	_, err := db.Exec("INSERT INTO users(email, password) VALUES($1, $2);", user.Email, user.Password)
	if err != nil {
		w.Write([]byte("Signup DB insert error: " + err.Error() + "\n"))
	} else {
		w.Write([]byte("Signup OK\n"))
	}
}

//configファイルの読み込み
func SetConfig() {
	if os.Getenv("AuthEnv") == "production" {
		fmt.Println("環境:production")
		viper.SetConfigName("config.production")
		viper.AddConfigPath(".")
		viper.ReadInConfig()
	} else {
		fmt.Println("環境:develop")
		viper.SetConfigName("config.develop")
		viper.AddConfigPath(".")
		viper.ReadInConfig()
	}
}
