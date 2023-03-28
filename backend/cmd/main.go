package main

import (
	"log"

	"diplomka/pkg/sqlite"
)

// const port = "8080"

func main() {
	db, err := sqlite.Connect("./test.db")
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	// m, err := migrate.New()

	// err = database.RunMigrateScripts(db)
	// if err != nil {
	// 	log.Fatalf("Run migration error: %v", err)
	// }

	// uRepo := repository.NewUserRepo(db)
	// sRepo := repository.NewSessionRepo(db)

	// authS := service.NewAuthService(uRepo, sRepo)

	// authH := handlers.NewAuthHandlers(authS)

	// r := mux.NewRouter()

	// r.HandleFunc("/login", authH.LogIn).Methods(http.MethodPost)

	// server := http.Server{
	// 	Addr:         ":" + port,
	// 	Handler:      r,
	// 	WriteTimeout: 5 * time.Second,
	// 	ReadTimeout:  5 * time.Second,
	// }

	// log.Printf("Starting listener on http://localhost:%s", port)

	// err = server.ListenAndServe()
	// log.Fatalf("Server error: %v", err)
}
