package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/florianwoelki/insta-clone/service.image-storage/files"
	"github.com/florianwoelki/insta-clone/service.image-storage/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var address = ":9090"
var basePath = "/tmp/images"
var env string

func main() {
	if env == "docker" {
		godotenv.Load(".env")
	} else {
		godotenv.Load("../.env")
	}

	logger := log.New(os.Stdout, "image-storage", log.LstdFlags)

	// create the storage class, use local storage
	// max filesize 5MB
	storage, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		logger.Fatal("Unable to create storage:", err)
		os.Exit(1)
	}

	// create the handlers
	fileHandler := handlers.NewFiles(storage, logger)

	// create gorilla mux router with CORS
	router := mux.NewRouter()

	postHandler := router.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fileHandler.ServeHTTP)

	// get files
	getHandler := router.Methods(http.MethodGet).Subrouter()
	getHandler.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))

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
