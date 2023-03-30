package main

import (
	"log"
	"net/http"
	"time"

	"diplomka/internal/handlers"
	"diplomka/internal/repository"
	"diplomka/internal/service"
	"diplomka/pkg/sqlite"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const port = "8080"

func main() {
	db, err := sqlite.Connect("./db/data.db")
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	uRepo := repository.NewUserRepo(db)
	sRepo := repository.NewSessionRepo(db)

	authS := service.NewAuthService(uRepo, sRepo)

	authH := handlers.NewAuthHandlers(authS)

	r := mux.NewRouter()

	r.HandleFunc("/login", authH.LogIn).Methods(http.MethodPost)
	r.HandleFunc("/signup", authH.SignUp).Methods(http.MethodPost)

	server := http.Server{
		Addr:         ":" + port,
		Handler:      r,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Printf("Starting listener on http://localhost:%s", port)

	err = server.ListenAndServe()
	log.Fatalf("Server error: %v", err)
}
