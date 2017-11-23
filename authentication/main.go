package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gaku3601/study-microservices/authentication/config"
	_ "github.com/lib/pq"

	mux "github.com/gorilla/mux.git"
)

func main() {
	config.SetConfig("./config")

	r := mux.NewRouter()
	r.HandleFunc("/users/login", Login).Methods("POST")
	r.HandleFunc("/users/signup", SignUp).Methods("POST")

	http.ListenAndServe(":8080", r)
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserTable struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
}

//ログイン認証。ログイン完了後、JWTトークンを返却する
func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	error := decoder.Decode(&user)
	if error != nil {
		w.Write([]byte("json decode error " + error.Error() + "\n"))
	}

	db, _ := sql.Open("postgres", "user=postgres host=localhost dbname=auth_db port=5433 sslmode=disable")
	defer db.Close()

	userTable := new(UserTable)
	err := db.QueryRow("SELECT id,email,password FROM users where email = $1;", user.Email).Scan(&userTable.ID, &userTable.Email, &userTable.Password)
	if err != nil {
		w.Write([]byte("emailが登録されていません。:" + err.Error() + "\n"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userTable.Password), []byte(user.Password))
	if err != nil {
		w.Write([]byte("email,passwordが違います。:" + err.Error() + "\n"))
		return
	}

	//認証処理
	res := new(Response)
	res.Token = fetchCreateToken()
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

	//passwordのhash化
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err := db.Exec("INSERT INTO users(email, password) VALUES($1, $2);", user.Email, bcryptPassword)
	if err != nil {
		w.Write([]byte("Signup DB insert error: " + err.Error() + "\n"))
	} else {
		w.Write([]byte("Signup OK\n"))
	}
}
