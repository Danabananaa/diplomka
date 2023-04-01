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

	// Добавленные новые репозитории

	spendingRepo := repository.NewSpendingTypeRepo(db)
	incometypeRepo := repository.NewIncomeTypeRepo(db)
	incomeRepo := repository.NewIncomeRepo(db)
	// test := repository.NewTestRepo(db)

	jwtService := service.NewJWTService()
	authService := service.NewAuthService(userRepo, jwtService)
	middlewareHandlers := handlers.NewMiddleware(authService)

	// Добавленные новые сервисы
	spendingTypeService := service.NewSpendingService(spendingRepo)
	incomeTypeService := service.NewIncomeTypeService(incometypeRepo)
	incomeService := service.NewIncomeService(incomeRepo)

	authHandlers := handlers.NewAuthHandlers(authService)
	// Добавленные новые хендлеры
	spendTH := handlers.NewSpendingHandlers(spendingTypeService)
	incTH := handlers.NewIncomeTypeHandlers(incomeTypeService)
	incH := handlers.NewIncomeHandlers(incomeService)

	r := mux.NewRouter()

	r.HandleFunc("/signup", authHandlers.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandlers.LogIn).Methods(http.MethodPost)

	// для проверки
	r.Handle("/", middlewareHandlers.RequireAuthentication(http.HandlerFunc(index))).Methods(http.MethodGet)

	// Добавленные новые хендлеры
	r.HandleFunc("/spending/type", spendTH.AllSpendingTypes).Methods(http.MethodGet)
	r.HandleFunc("/income/type", incTH.AllIncomeTypes).Methods(http.MethodGet)
	r.HandleFunc("/income/", incH.PostIncome).Methods(http.MethodPost)

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
