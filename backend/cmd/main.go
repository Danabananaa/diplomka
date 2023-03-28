package main

import (
	"net/http"
	"time"

	"diplomka/internal/handlers"
	"diplomka/internal/repository"
	"diplomka/internal/service"
	"diplomka/pkg/log"
	"diplomka/pkg/sqlite"

	"github.com/gorilla/mux"
)

const port = "8080"

func main() {
	db, err := sqlite.Connect("./db/data.db")
	if err != nil {
		log.Fatalln(db)
	}

	uRepo := repository.NewUserRepo(db)
	sRepo := repository.NewSessionRepo(db)

	authS := service.NewAuthService(uRepo, sRepo)

	authH := handlers.NewAuthHandlers(authS)

	r := mux.NewRouter()

	r.HandleFunc("/login", authH.LogIn).Methods(http.MethodPost)

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
