package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/florianwoelki/insta-clone/account-api/handlers"
	"github.com/florianwoelki/insta-clone/internal"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var address = ":9090"

func main() {
	godotenv.Load("../.env") // TODO: change path for docker env

	logger := log.New(os.Stdout, "accounts-api ", log.LstdFlags)
	validator := internal.NewValidation()

	// create connection to postgres database
	database, gormInstance := internal.NewDatabase(logger).CreateConnection()
	defer database.Close()

	// create gorilla mux router
	router := mux.NewRouter()

	// create handlers
	accountsHandler := handlers.NewAccounts(logger, gormInstance, validator)

	// handlers for API
	getAccounts := router.Methods(http.MethodGet).Subrouter()
	getAccounts.HandleFunc("/account/{id:[0-9]+}", accountsHandler.ListSingle)

	postAccounts := router.Methods(http.MethodPost).Subrouter()
	postAccounts.HandleFunc("/login", accountsHandler.Login)
	postAccounts.HandleFunc("/refreshToken", accountsHandler.Refresh)

	registerAccounts := router.Methods(http.MethodPost).Subrouter()
	registerAccounts.HandleFunc("/register", accountsHandler.Register)
	registerAccounts.Use(accountsHandler.MiddlewareValidateAccount)

	// create a new server
	server := http.Server{
		Addr:         address,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // keeps the connection open for specified time
	}

	// start the server
	go func() {
		logger.Println("Starting server on port", address)

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// gracefully shutdown the server and catch interrupt or kill
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until signal is received
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
