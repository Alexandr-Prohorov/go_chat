package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewClaims(username string) *Claims {
	expirationTime := time.Now().Add(5 * time.Minute)
	return &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
}

var jwtKey = []byte("my_secret_key")

func (v *Claims) GenerateJwt() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, v)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil
	}

	return tokenString, err

}
