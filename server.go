package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/Muchogoc/semezana/adapters"
	"github.com/Muchogoc/semezana/app"
	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/domain/user"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

var (
	DB_HOST     string = os.Getenv("DB_HOST")
	DB_PORT, _         = strconv.Atoi(os.Getenv("DB_PORT"))
	DB_USER     string = os.Getenv("DB_USER")
	DB_NAME     string = os.Getenv("DB_NAME")
	DB_PASSWORD string = os.Getenv("DB_PASSWORD")
)

var (
	NSQ_ADDRESS        string = os.Getenv("NSQ_ADDRESS")
	NSQ_LOOKUP_ADDRESS string = os.Getenv("NSQ_LOOKUP_ADDRESS")
)

var (
	PORT        = os.Getenv("PORT")
	DEBUG, _    = strconv.ParseBool(os.Getenv("DEBUG"))
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)

func loadSampleData(db *sql.DB) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Paths(
			"data/dev",
		),
	)
	if err != nil {
		return err
	}

	if err := fixtures.Load(); err != nil {
		return err
	}

	return nil
}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()

	client := ent.NewClient(
		ent.Driver(entsql.OpenDB("postgres", db)),
	)
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if ENVIRONMENT == "local" {
		if err := loadSampleData(db); err != nil {
			log.Fatalf("failed to load sample data: %v", err)
		}
	}

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(NSQ_ADDRESS, config)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Stop()

	if DEBUG {
		producer.SetLogger(log.Default(), nsq.LogLevelDebug)
		client = client.Debug()
	}

	chatFactory, err := chat.NewFactory()
	if err != nil {
		log.Fatal(err)
	}

	userFactory, err := user.NewFactory()
	if err != nil {
		log.Fatal(err)
	}

	repository := adapters.NewEntRepository(client, chatFactory, userFactory)
	publisher := adapters.NewNSQPublisher(producer)
	subscriber := adapters.NewNSQSubscriber(NSQ_LOOKUP_ADDRESS)

	chatService := app.NewChatService(repository, repository, publisher, subscriber)
	sessionStore := ports.NewSessionStore()
	defer sessionStore.Shutdown()

	mux := service(chatService, sessionStore)

	server := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%s", PORT),
	}

	go func() {
		logrus.Infof("Server running at port :%s", PORT)
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
	// r.Use(middleware.Logger)

	// r.Use(middleware.Heartbeat("/ping"))
	// logger := httplog.NewLogger("httplog", httplog.Options{
	// 	JSON:    !DEBUG,
	// 	Concise: DEBUG,
	// })
	// r.Use(httplog.RequestLogger(logger))

	// r.Use(middleware.Recoverer)

	// r.Mount("/debug", middleware.Profiler())

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
