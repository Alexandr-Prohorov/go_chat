package models

type Message struct {
	MessageId int
	ChatId    int
	SenderId  int
	Content   string `json:"Content"`
	CreatedAt string
	UpdatedAt string
}
