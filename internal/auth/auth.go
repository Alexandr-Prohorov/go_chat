package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Login    string
	Password string
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	a := []byte(password)
	fmt.Println(a)

	return string(hash), err
}
