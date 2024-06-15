package api

import (
	"chat-app/internal/utils"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"chat-app/internal/models"
	"chat-app/internal/store"
)

type UserHandler struct {
	Store store.UserStore
}

func NewUserHandler(store store.UserStore) *UserHandler {
	return &UserHandler{
		Store: store,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	//username := r.Context().Value("username").(string)
	//fmt.Println(username)
	//token := middleware.AuthMiddleware(h.JWTSecretKey)
	//
	//fmt.Println(token)

	//username := r.Context().Value("username").(string)
	//if username == "" {
	//	http.Error(w, "Username not found", http.StatusUnauthorized)
	//	return
	//}

	username := utils.GetContextValue(w, r)

	users, err := h.Store.GetAllUsers(username)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetOneUser(w http.ResponseWriter, r *http.Request) {
	username := utils.GetContextValue(w, r)

	users, err := h.Store.GetOneUser(username)
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil { // сравнивает модель данных с телом запроса
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	if err := h.Store.AddUser(&user, hashedPassword); err != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
