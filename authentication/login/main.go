package login

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gaku3601/study-microservices/authentication/dbc"
	"golang.org/x/crypto/bcrypt"
)

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

	dbc.DBConnect(func(db *sql.DB) {
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
	})
}
