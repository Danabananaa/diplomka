package main

import (
	"diplomka/internal/handlers"
	"diplomka/internal/repository"
	"diplomka/internal/service"
	"diplomka/pkg/sqlite"
	"fmt"
	"log"
	"net/http"
	"time"

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
	userRepo := repository.NewUserRepo(db)
	spendingRepo := repository.NewSpendingTypeRepo(db)
	incomeRepo := repository.NewIncomeTypeRepo(db)

	jwtService := service.NewJWTService()

	authService := service.NewAuthService(userRepo, jwtService)
	spendingService := service.NewSpendingService(spendingRepo)
	incomeService := service.NewIncomeService(incomeRepo)
	middlewareHandlers := handlers.NewMiddleware(authService)
	authHandlers := handlers.NewAuthHandlers(authService)
	sH := handlers.NewSpendingHandlers(spendingService)
	iH := handlers.NewIncomeHandlers(incomeService)

	r := mux.NewRouter()

	r.HandleFunc("/signup", authHandlers.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandlers.LogIn).Methods(http.MethodPost)

	// для проверки
	r.Handle("/", middlewareHandlers.RequireAuthentication(http.HandlerFunc(index))).Methods(http.MethodGet)
	r.HandleFunc("/spending/type", sH.AllSpendingTypes).Methods(http.MethodGet)
	r.HandleFunc("/income/type", iH.AllIncomeTypes).Methods(http.MethodGet)

	r.Use(middlewareHandlers.PanicRecover)

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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "work")
}
