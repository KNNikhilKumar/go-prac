package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Name     string
	LoginId  string
	Password string
}

type Claims struct {
	Name    string
	LoginId string
	jwt.MapClaims
}

func StartHandlers() {
	http.HandleFunc("/create", createToken)
	http.HandleFunc("/authenticate", authToken)

}

func createToken(w http.ResponseWriter, r *http.Request) {
}

func authToken(w http.ResponseWriter, r *http.Request) {

}
