package api

import (
	"chat-app/internal/store"
	"chat-app/internal/utils"
	"encoding/json"
	"net/http"
)

type ChatHandler struct {
	Store store.ChatStore
}

func NewChatHandler(store store.ChatStore) *ChatHandler {
	return &ChatHandler{
		Store: store,
	}
}

func (h *ChatHandler) GetOneChat(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetContextValue(w, r, "id")

	users, err := h.Store.GetOneChat(userId)
	if err != nil {
		http.Error(w, "Failed to fetch chat", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
