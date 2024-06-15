package store

import (
	"database/sql"

	"chat-app/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) GetAllUsers(user string) ([]*models.User, error) { // метод экземпляра UserStore, возврващает
	var users []*models.User
	rows, err := s.db.Query("SELECT id, username, email FROM users WHERE username != $1", user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (s *UserStore) GetOneUser(user string) (*models.User, error) {
	var oneUser models.User

	row := s.db.QueryRow("SELECT id, username, email FROM users WHERE username = $1", user)

	if err := row.Scan(&oneUser.ID, &oneUser.Username, &oneUser.Email); err != nil {
		return nil, err
	}

	return &oneUser, nil
}

func (s *UserStore) AddUser(user *models.User, hashedPassword []byte) error {
	_, err := s.db.Exec("INSERT INTO users (username, surname, email, password) VALUES ($1, $2, $3, $4)", user.Username, user.Surname, user.Email, hashedPassword)
	return err
}
