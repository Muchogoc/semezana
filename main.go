package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Muchogoc/semezana/semezana"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	stop := make(chan bool)
	done := make(chan bool)
	addr := ":8080"

	mux := semezana.PrepareServer(ctx, stop, done)

	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}

	go func() {
		logrus.Infof("Server running at port %v", addr)
		if err := server.ListenAndServe(); err != nil {
			logrus.Println("HTTP server failed to listen and server", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-signalChan

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		logrus.Println("HTTP server failed to terminate gracefully", err)
	}

	<-ctx.Done()

	stop <- true

	<-done
}
