package main

import (
	"chat-app/internal/api"
	"chat-app/internal/middleware"
	"chat-app/internal/store"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:157947XXxx@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err) // логируем если соединение не установлено
	}
	defer db.Close() // гарантия закрытия соединения

	if err = db.Ping(); err != nil { // проверка соединения
		log.Fatal(err)
	}

	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	fmt.Println(file)

	var config map[string]string
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
	}

	jwtKey := config["jwt_secret_key"]

	fmt.Println(jwtKey)

	userStore := store.NewUserStore(db)
	authStore := store.NewAuthStore(db)
	userHandler := api.NewUserHandler(*userStore)
	authHandler := api.NewAuthHandler(*authStore)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/create", userHandler.AddUser).Methods("POST")
	r.HandleFunc("/", getUi)
	r.HandleFunc("/auth", authHandler.Auth).Methods("POST")

	r.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("ui", "styles.css")
		file, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Could not read file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/css")
		w.Write(file)
	})

	r.HandleFunc("/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("ui", "scripts.js")
		file, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Could not read file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(file)
	})

	r.HandleFunc("/", getUi) // TODO: переделать
	r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims := &middleware.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		filePath := filepath.Join("ui", "main.html")
		file, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Could not read file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(file)
	}) // TODO: переделать

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUi(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("ui", "index.html")
	file, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(file)
}
