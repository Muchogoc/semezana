package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Muchogoc/semezana/ent"
	_ "github.com/mattn/go-sqlite3"
)

var globals struct {
	maxMessageSize int64
	sessionStore   *SessionStore
	hub            *Hub
	client         *ent.Client
}

func main() {
	mux := http.NewServeMux()

	client, err := ent.Open("sqlite3", "file:semezana.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	globals.sessionStore = NewSessionStore()

	// The hub (the main message router)
	globals.hub = newHub()

	globals.client = client.Debug()

	mux.HandleFunc("/channels", serveWebSocket)

	if err = listenAndServe(":8080", mux, signalHandler()); err != nil {
		log.Fatal(err)
	}

}

func signalHandler() <-chan bool {
	stop := make(chan bool)

	signchan := make(chan os.Signal, 1)
	signal.Notify(signchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		// Wait for a signal. Don't care which signal it is
		sig := <-signchan
		log.Printf("Signal received: '%s', shutting down", sig)
		stop <- true
	}()

	return stop
}
