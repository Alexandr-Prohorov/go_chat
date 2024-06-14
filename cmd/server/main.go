package main

import (
	"chat-app/internal/api"
	"chat-app/internal/store"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type Config struct {
	DBUser       string `json:"db_user"`
	DBPassword   string `json:"db_password"`
	DBName       string `json:"db_name"`
	DBHost       string `json:"db_host"`
	DBPort       string `json:"db_port"`
	JWTSecretKey string `json:"jwt_secret_key"`
}

func main() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
	}

	// Формирование строки подключения
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	// Подключение к базе данных
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // гарантия закрытия соединения

	if err = db.Ping(); err != nil { // проверка соединения
		log.Fatal(err)
	}

	userStore := store.NewUserStore(db)
	authStore := store.NewAuthStore(db)
	userHandler := api.NewUserHandler(*userStore)
	authHandler := api.NewAuthHandler(*authStore)

	indexFile := api.NewStaticFile("/views/index.html", "text/html")
	mainFile := api.NewStaticFile("/views/main.html", "text/html")
	stylesFile := api.NewStaticFile("/styles.css", "text/css")
	scriptsFile := api.NewStaticFile("/scripts/scripts.js", "text/javascript")
	mainScriptsFile := api.NewStaticFile("/scripts/main_script.js", "text/javascript")

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/create", userHandler.AddUser).Methods("POST")
	r.HandleFunc("/auth", authHandler.Auth).Methods("POST")

	r.HandleFunc("/", indexFile.StaticFileHandler)
	r.HandleFunc("/main", mainFile.StaticFileHandler)
	r.HandleFunc("/scripts.js", scriptsFile.StaticFileHandler)
	r.HandleFunc("/main_script.js", mainScriptsFile.StaticFileHandler)
	r.HandleFunc("/styles.css", stylesFile.StaticFileHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
