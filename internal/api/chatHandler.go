package api

import (
	"chat-app/internal/models"
	"chat-app/internal/store"
	"chat-app/internal/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	Store store.ChatStore
}

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewChatHandler(store store.ChatStore) *ChatHandler {
	return &ChatHandler{
		Store: store,
	}
}

func (h *ChatHandler) GetOneChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // не есть правильно, использую разные способы получения id

	userId := utils.GetContextValue(w, r, "id")
	memberId := vars["id"]

	users, err := h.Store.GetChats(userId, memberId)
	if err != nil {
		http.Error(w, "Failed to fetch chat", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *ChatHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chatId := vars["id"]

	users, err := h.Store.GetMessages(chatId)
	if err != nil {
		http.Error(w, "Failed to fetch chat", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *ChatHandler) ChatRoomWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID, err := strconv.Atoi(vars["chat_id"])
	userId := utils.GetContextValue(w, r, "id")
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	// Обновляем соединение до WebSocket
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Обрабатываем WebSocket соединение
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received message in chat %d: %s", chatID, message)
		var content models.Message
		err = json.Unmarshal([]byte(message), &content)
		// Пример обработки входящего сообщения (например, сохранение в БД)
		if err := h.Store.AddMessage(strconv.Itoa(chatID), userId, content.Content); err != nil {
			http.Error(w, "Failed to add message", http.StatusInternalServerError)
			return
		}

		// Пример отправки ответа всем подключенным клиентам (здесь мы просто отправляем обратно тому же клиенту)
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
