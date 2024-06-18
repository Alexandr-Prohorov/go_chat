package store

import (
	"chat-app/internal/models"
	"database/sql"
)

type ChatStore struct {
	db *sql.DB
}

func NewChatStore(db *sql.DB) *ChatStore {
	return &ChatStore{db: db}
}

func (s *ChatStore) GetOneChat(userId string) (*models.Chat, error) {
	var oneChat models.Chat

	row := s.db.QueryRow("SELECT * FROM chats WHERE owner_id = $1", userId)

	if err := row.Scan(
		&oneChat.ChatId,
		&oneChat.Name,
		&oneChat.Description,
		&oneChat.OwnerId,
		&oneChat.CreatedAt,
		&oneChat.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &oneChat, nil
}
