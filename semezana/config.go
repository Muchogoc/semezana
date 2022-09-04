package semezana

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Muchogoc/semezana/ent"
	_ "github.com/lib/pq"
	"github.com/nsqio/go-nsq"
)

var globals struct {
	maxMessageSize int64
	sessionStore   *SessionStore
	client         *ent.Client
	producer       *nsq.Producer
}

var DEBUG = false

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

func PrepareServer(ctx context.Context, stop <-chan bool, done chan<- bool) *http.ServeMux {
	mux := http.NewServeMux()

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

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//The only valid way to instantiate the Config
	config := nsq.NewConfig()

	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer(NSQ_ADDRESS, config)
	if err != nil {
		log.Fatal(err)
	}

	if DEBUG {
		producer.SetLogger(log.Default(), nsq.LogLevelDebug)
		client.Debug()
	}

	globals.sessionStore = NewSessionStore()

	globals.client = client

	globals.producer = producer

	mux.HandleFunc("/socket", serveWebSocket)

	go func() {
		<-stop

		client.Close()
		producer.Stop()
		globals.sessionStore.Shutdown()

		done <- true
	}()

	return mux
}
