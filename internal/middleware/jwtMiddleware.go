package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Jwt struct {
	Key string `json:"jwt_secret_key"`
}

func NewClaims(username string) *Claims {
	//expirationTime := time.Now().Add(5 * time.Hour)
	return &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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
		fmt.Println("Error unmarshaling file:", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, v)
	tokenString, err := token.SignedString([]byte(jwtKey.Key))
	if err != nil {
		return "", nil
	}

	return tokenString, err

}
