package main

import (
	"diplomka/internal/handlers"
	"diplomka/internal/repository"
	"diplomka/internal/service"
	"diplomka/pkg/sqlite"
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
	jwtService := service.NewJWTService()

	userRepo := repository.NewUserRepo(db)
	spendingTypeRepo := repository.NewSpendingTypeRepo(db)
	incometypeRepo := repository.NewIncomeTypeRepo(db)
	incomeRepo := repository.NewIncomeRepo(db)
	spendingRepo := repository.NewSpendingRepo(db)
	ldtyperepo := repository.NewLoanDebtTypeRepo(db)
	loanRepo := repository.NewLoanRepo(db)
	debtRepo := repository.NewDebtRepo(db)
	statRepo := repository.NewStatisticsRepo(db)

	// services
	authService := service.NewAuthService(userRepo, jwtService)

	incSpnTypeService := service.NewIncSpnTypeService(incometypeRepo, spendingTypeRepo)
	incomeService := service.NewIncomeService(incomeRepo)
	spendingService := service.NewSpendingService(spendingRepo)
	ldtypeService := service.NewLoanDebtTypeService(ldtyperepo)
	loanDebtService := service.NewLoanDebtService(loanRepo, debtRepo)
	statService := service.NewStatisticsService(statRepo)

	// handlers
	mHandlers := handlers.NewMiddleware(authService)

	authHandlers := handlers.NewAuthHandlers(authService)

	incSpnTypeHandlers := handlers.NewIncomeSpendTypeHandlers(incSpnTypeService)
	incomeHandlers := handlers.NewIncomeHandlers(incomeService)
	budgetHandler := handlers.NewIncomeSpendingHandlers(incomeService, spendingService)
	spendingHadlers := handlers.NewSpendingHandlers(spendingService)
	ldtypeHandlers := handlers.NewLoanDebtTypeHandlers(ldtypeService)
	loanDebtHandlers := handlers.NewLoanDebtHandlers(loanDebtService)
	statHandlers := handlers.NewStatisticsHandlers(statService)

	r := mux.NewRouter()

	// r.Handle("/", middlewareHandlers.RequireAuthentication(http.HandlerFunc(index))).Methods(http.MethodGet)

	r.HandleFunc("/signup", authHandlers.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandlers.LogIn).Methods(http.MethodPost)
	// Added new hadler which contain getting income and spending
	r.HandleFunc("/budget/type", incSpnTypeHandlers.AllIncomeTypes).Methods(http.MethodGet)
	// Added new hadler which contain getting assets and liability types
	r.HandleFunc("/debt/type", ldtypeHandlers.GetAllLoanDebtType).Methods(http.MethodGet)
	r.Handle("/userinfo", mHandlers.RequireAuthentication(http.HandlerFunc(authHandlers.GetUserInfo))).Methods(http.MethodGet)
	r.Handle("/userinfo", mHandlers.RequireAuthentication(http.HandlerFunc(authHandlers.PostImage))).Methods(http.MethodPost)
	r.Handle("/image", mHandlers.RequireAuthentication(http.HandlerFunc(authHandlers.GetImage))).Methods(http.MethodGet)

	r.Handle("/spending", mHandlers.RequireAuthentication(http.HandlerFunc(spendingHadlers.PostSpending))).Methods(http.MethodPost)
	r.Handle("/income", mHandlers.RequireAuthentication(http.HandlerFunc(incomeHandlers.PostIncome))).Methods(http.MethodPost)

	r.Handle("/debt", mHandlers.RequireAuthentication(http.HandlerFunc(loanDebtHandlers.AddLoanDebt))).Methods(http.MethodPost)
	r.Handle("/debt", mHandlers.RequireAuthentication(http.HandlerFunc(loanDebtHandlers.GetLoanDebt))).Methods(http.MethodGet)

	r.Handle("/budget/stats", mHandlers.RequireAuthentication(http.HandlerFunc(budgetHandler.GetIncomeSpending))).Methods(http.MethodGet)

	r.Handle("/statistics", mHandlers.RequireAuthentication(http.HandlerFunc(statHandlers.GetStat))).Methods(http.MethodGet)

	r.Use(mHandlers.PanicRecover)

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
