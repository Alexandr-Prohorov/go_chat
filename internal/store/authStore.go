package store

import (
	"chat-app/internal/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) *AuthStore {
	return &AuthStore{db: db}
}

func (s *AuthStore) GetUser(u *models.Auth) (*models.User, error) {
	var user models.User

	row := s.db.QueryRow("SELECT * FROM users WHERE username = $1", u.Login)

	err := row.Scan(&user.ID, &user.Username, &user.Surname, &user.Surname, &user.Password)
	if err != nil {
		return nil, err
	}

	hashedPassword := []byte(user.Password)

	// Сравнение пароля
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(u.Password))
	if err != nil {
		// Пароль неверный
		return nil, fmt.Errorf("invalid password")
	}

	return &user, err
}
