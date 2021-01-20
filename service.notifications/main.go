package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var address = ":9090"
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(connection *websocket.Conn) {
	for {
		messageType, messageBytes, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(messageBytes))

		if err := connection.WriteMessage(messageType, messageBytes); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// allow any connection to ws endpoint... TODO: fix
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client successfully connected to WebSocket...")
	reader(websocket)
}

func main() {
	logger := log.New(os.Stdout, "notifications ", log.LstdFlags)

	router := mux.NewRouter()
	router.HandleFunc("/", wsEndpoint)

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
