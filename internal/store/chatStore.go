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

func (s *ChatStore) GetChats(userId string, memberId string) ([]*models.Chat, error) {
	var chats []*models.Chat

	rows, err := s.db.Query("SELECT c.chat_id FROM chats c JOIN chat_members cm1 ON c.chat_id = cm1.chat_id JOIN chat_members cm2 ON c.chat_id = cm2.chat_id WHERE cm1.user_id = 	$1 AND  cm2.user_id = $2;", userId, memberId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var oneChat models.Chat
		if err := rows.Scan(&oneChat.ChatId); err != nil {
			return nil, err
		}
		chats = append(chats, &oneChat)
	}

	return chats, nil
}

func (s *ChatStore) GetMessages(chatId string) ([]*models.Message, error) {
	var messages []*models.Message

	rows, err := s.db.Query("SELECT * FROM messages WHERE chat_id = $1;", chatId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.MessageId, &message.ChatId, &message.SenderId, &message.Content, &message.CreatedAt, &message.UpdatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

func (s *ChatStore) AddMessage(chatId string, userId string, content string) error {
	_, err := s.db.Exec("INSERT INTO messages (chat_id, sender_id, content) VALUES ($1, $2, $3)", chatId, userId, content)
	return err
}
