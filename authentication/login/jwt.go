package login

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	viper "github.com/spf13/viper.git"
)

type JwtKey struct {
	Created_at  int    `json:"created_at"`
	Id          string `json:"id"`
	Algorithm   string `json:"algorithm"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	Consumer_id string `json:"consumer_id"`
	UserID      string
}

func fetchCreateToken(userID string) string {
	url := viper.GetString("TokenKeyServerURL")
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

	jwtKey.UserID = userID
	defer resp.Body.Close()

	return createTokenString(jwtKey)
}

func createTokenString(jwtKey *JwtKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  jwtKey.UserID,
		"iss": jwtKey.Key,
	})
	tokenString, _ := token.SignedString([]byte(jwtKey.Secret))

	return tokenString
}
