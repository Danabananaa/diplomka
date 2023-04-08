package main

import (
	"log"
	"net/http"
	"time"

	"diplomka/internal/handlers"
	"diplomka/internal/repository"
	"diplomka/internal/service"
	"diplomka/pkg/sqlite"

	middleware "diplomka/internal/handlers/handlers_middleware"

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

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		log.Fatalln(err)
	}

	repo := repository.NewRepository(db)

	authService := service.NewAuthenticationService(repo)
	finansService := service.NewFinancialTrackerService(repo)
	avatarService := service.NewAvatarService(repo)

	r := mux.NewRouter()

	m := middleware.NewMiddleware(authService)

	handlers.InitAuthHandlers(r, authService)
	handlers.InitFinancialHandlers(r, m, finansService)
	handlers.InitAvatarHandlers(r, m, avatarService)

	r.Use(m.PanicRecover)

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
