package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Muchogoc/semezana/adapters"
	"github.com/Muchogoc/semezana/app"
	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

var (
	DB_HOST     string = os.Getenv("DB_HOST")
	DB_PORT     string = os.Getenv("DB_PORT")
	DB_USER     string = os.Getenv("DB_USER")
	DB_NAME     string = os.Getenv("DB_NAME")
	DB_PASSWORD string = os.Getenv("DB_PASSWORD")
)

var (
	NSQ_ADDRESS        string = os.Getenv("NSQ_ADDRESS")
	NSQ_LOOKUP_ADDRESS string = os.Getenv("NSQ_LOOKUP_ADDRESS")
)

var (
	PORT  = "8080"
	DEBUG = false
)

func main() {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME,
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(NSQ_ADDRESS, config)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Stop()

	if DEBUG {
		producer.SetLogger(log.Default(), nsq.LogLevelDebug)
		client.Debug()
	}

	chatFactory, err := chat.NewFactory()
	if err != nil {
		log.Fatal(err)
	}

	db := adapters.NewEntRepository(client, chatFactory)
	publisher := adapters.NewNSQPublisher(producer)
	subscriber := adapters.NewNSQSubscriber(NSQ_LOOKUP_ADDRESS)

	chatService := app.NewChatService(db, db, publisher, subscriber)
	sessionStore := ports.NewSessionStore()
	defer sessionStore.Shutdown()

	mux := service(chatService, sessionStore)

	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", PORT),
	}

	go func() {
		logrus.Info("Server running at port %v", PORT)
		if err := server.ListenAndServe(); err != nil {
			logrus.Println("HTTP server failed to listen and server", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-signalChan

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logrus.Println("HTTP server failed to terminate gracefully", err)
	}

	<-ctx.Done()
}

func service(service app.ChatService, store *ports.SessionStore) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	addCorsMiddleware(r)

	r.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	r.Use(middleware.NoCache)

	ports.HandlerFromMux(ports.NewHttpServer(service, store), r)

	return r
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(corsMiddleware.Handler)
}
