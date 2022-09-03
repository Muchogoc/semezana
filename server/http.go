package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

func listenAndServe(addr string, mux *http.ServeMux, stop <-chan bool) error {
	httpdone := make(chan bool)

	server := &http.Server{
		Handler: mux,
	}

	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Println("HTTP listener: failed", err)
		} else {
			err := server.Serve(lis)
			if err != nil {
				log.Println("HTTP server: failed", err)
			}
		}

		httpdone <- true
	}()

Loop:
	for {
		select {
		case <-stop:
			// Give server 2 seconds to shut down.
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			if err := server.Shutdown(ctx); err != nil {
				// failure/timeout shutting down the server gracefully
				log.Println("HTTP server failed to terminate gracefully", err)
			}

			// While the server shuts down, terminate all sessions.
			globals.sessionStore.Shutdown()

			// Wait for http server to stop Accept()-ing connections.
			<-httpdone
			cancel()

			break Loop

		case <-httpdone:
			break Loop
		}
	}

	return nil
}
