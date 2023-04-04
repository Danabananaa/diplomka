package main

import (
	"fmt"
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

	// repository
	userRepo := repository.NewUserRepo(db)
	spendingTypeRepo := repository.NewSpendingTypeRepo(db)
	incometypeRepo := repository.NewIncomeTypeRepo(db)
	incomeRepo := repository.NewIncomeRepo(db)
	spendingRepo := repository.NewSpendingRepo(db)

	jwtService := service.NewJWTService()
	authService := service.NewAuthService(userRepo, jwtService)

	// services
	spendingTypeService := service.NewSpendingTypeService(spendingTypeRepo)
	incomeTypeService := service.NewIncomeTypeService(incometypeRepo)
	incomeService := service.NewIncomeService(incomeRepo)
	spendingService := service.NewSpendingService(spendingRepo)

	// handlers
	middlewareHandlers := handlers.NewMiddleware(authService)

	authHandlers := handlers.NewAuthHandlers(authService)

	spendingTypeHandlers := handlers.NewSpendingTypeHandlers(spendingTypeService)
	incomeTypeHandlers := handlers.NewIncomeTypeHandlers(incomeTypeService)
	incomeHandlers := handlers.NewIncomeHandlers(incomeService)
	spendingHadlers := handlers.NewSpendingHandlers(spendingService)

	r := mux.NewRouter()

	// r.Handle("/", middlewareHandlers.RequireAuthentication(http.HandlerFunc(index))).Methods(http.MethodGet)

	r.HandleFunc("/signup", authHandlers.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandlers.LogIn).Methods(http.MethodPost)

	r.HandleFunc("/spending/type", spendingTypeHandlers.AllSpendingTypes).Methods(http.MethodGet)
	r.HandleFunc("/income/type", incomeTypeHandlers.AllIncomeTypes).Methods(http.MethodGet)

	r.Handle("/spending", middlewareHandlers.RequireAuthentication(http.HandlerFunc(spendingHadlers.PostSpending))).Methods(http.MethodPost)
	r.Handle("/income", middlewareHandlers.RequireAuthentication(http.HandlerFunc(incomeHandlers.PostIncome))).Methods(http.MethodPost)

	// r.HandleFunc("/spending", incomeHandlers.PostIncome).Methods(http.MethodGet)
	r.Handle("/income", middlewareHandlers.RequireAuthentication(http.HandlerFunc(incomeHandlers.GetIncome))).Methods(http.MethodGet)

	r.HandleFunc("/statistics", incomeHandlers.PostIncome).Methods(http.MethodGet)

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

// type statistics struct {
// 	start string
// 	end   string
// }
// type name struct {
// 	Amount float64
// 	Type   model.SpendingType
// }
// type name2 struct {
// 	Amount float64
// 	Type   model.SpendingType
// }

// type statistics struct {
// 	s []name
// 	i []name2
// }

func Statistics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "work")
}
