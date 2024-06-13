package api

import (
	"chat-app/internal/middleware"
	"chat-app/internal/models"
	"chat-app/internal/store"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Store store.AuthStore
}

func NewAuthHandler(store store.AuthStore) *AuthHandler {
	return &AuthHandler{
		Store: store,
	}
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var auth models.Auth

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil { // сравнивает модель данных с телом запроса
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	user, err := h.Store.GetUser(&auth)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	claims := middleware.NewClaims(user.Username)

	token, err := claims.GenerateJwt()
	if err != nil {
		http.Error(w, "ХЗчо указатьб", http.StatusForbidden)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Token",
		Value: token,
		Path:  "/",
	})

	w.WriteHeader(http.StatusCreated)

}
