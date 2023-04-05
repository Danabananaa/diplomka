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

	// repository
	userRepo := repository.NewUserRepo(db)
	spendingTypeRepo := repository.NewSpendingTypeRepo(db)
	incometypeRepo := repository.NewIncomeTypeRepo(db)
	incomeRepo := repository.NewIncomeRepo(db)
	spendingRepo := repository.NewSpendingRepo(db)
	assliatype := repository.NewAssLiaTypeRepo(db)
	assetsRepo := repository.NewAssetsRepo(db)
	liabsRepo := repository.NewLiabilitiesRepo(db)

	jwtService := service.NewJWTService()
	authService := service.NewAuthService(userRepo, jwtService)

	// services

	incSpnTypeService := service.NewIncSpnTypeService(incometypeRepo, spendingTypeRepo)
	incomeService := service.NewIncomeService(incomeRepo)
	spendingService := service.NewSpendingService(spendingRepo)
	assliatypeService := service.NewAssLiaTypeService(assliatype)
	assets_liab_Service := service.NewAssetsService(assetsRepo, liabsRepo)

	// handlers
	middlewareHandlers := handlers.NewMiddleware(authService)

	authHandlers := handlers.NewAuthHandlers(authService)

	incSpnTypeHandlers := handlers.NewIncomeSpendTypeHandlers(incSpnTypeService)
	incomeHandlers := handlers.NewIncomeHandlers(incomeService)
	budgetHandler := handlers.NewIncomeSpendingHandlers(incomeService, spendingService)
	spendingHadlers := handlers.NewSpendingHandlers(spendingService)
	assliatypeHandlers := handlers.NewAssLiaTypeHandlers(assliatypeService)
	assets_liab_Handlers := handlers.NewAssetsHandlers(assets_liab_Service)

	r := mux.NewRouter()

	// r.Handle("/", middlewareHandlers.RequireAuthentication(http.HandlerFunc(index))).Methods(http.MethodGet)

	r.HandleFunc("/signup", authHandlers.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandlers.LogIn).Methods(http.MethodPost)
	// Added new hadler which contain getting income and spending
	r.HandleFunc("/budget/type", incSpnTypeHandlers.AllIncomeTypes).Methods(http.MethodGet)
	// Added new hadler which contain getting assets and liability types
	r.HandleFunc("/assliatype", assliatypeHandlers.GetAllAssLiaType).Methods(http.MethodGet)

	r.Handle("/spending", middlewareHandlers.RequireAuthentication(http.HandlerFunc(spendingHadlers.PostSpending))).Methods(http.MethodPost)
	r.Handle("/income", middlewareHandlers.RequireAuthentication(http.HandlerFunc(incomeHandlers.PostIncome))).Methods(http.MethodPost)
	// Added new hadler which contain getting adding and getting assets and liability
	r.Handle("/debt", middlewareHandlers.RequireAuthentication(http.HandlerFunc(assets_liab_Handlers.AddAssetsLiabs))).Methods(http.MethodPost)
	r.Handle("/debt", middlewareHandlers.RequireAuthentication(http.HandlerFunc(assets_liab_Handlers.GetAssetLiab))).Methods(http.MethodGet)
	// Added new hadler which contain getting income and spending
	r.Handle("/budget/stats", middlewareHandlers.RequireAuthentication(http.HandlerFunc(budgetHandler.GetIncomeSpending))).Methods(http.MethodGet)

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

// {
//     "start_date":"2023-01-01",
//     "end_date":"2023-04-05"
// }
