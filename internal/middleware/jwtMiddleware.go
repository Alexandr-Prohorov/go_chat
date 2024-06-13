package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Jwt struct {
	Key string `json:"jwt_secret_key"`
}

func NewClaims(username string) *Claims {
	expirationTime := time.Now().Add(5 * time.Hour)
	return &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
}

func (v *Claims) GenerateJwt() (string, error) {
	file, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatal("Error parsing config file:", err)
		return "", err
	}

	var jwtKey Jwt
	err = json.Unmarshal(file, &jwtKey)
	if err != nil {
		fmt.Println("Error unmarshaling jwtKey:", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, v)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil
	}

	return tokenString, err

}
