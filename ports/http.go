package ports

import (
	"log"
	"net/http"

	"github.com/Muchogoc/semezana/app"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
)

type HttpServer struct {
	service      app.ChatService
	sessionStore *SessionStore
}

func NewHttpServer(service app.ChatService, sessionStore *SessionStore) HttpServer {
	return HttpServer{
		service:      service,
		sessionStore: sessionStore,
	}
}

// Create a channel
func (h HttpServer) CreateChannel(w http.ResponseWriter, r *http.Request) {
	newChannel := NewChannel{}
	if err := render.Decode(r, &newChannel); err != nil {
		render.Render(w, r, nil)
		return
	}
}

// Delete a channel
func (h HttpServer) DeleteChannelById(w http.ResponseWriter, r *http.Request, channelID string) {}

// Retrieve a channel
func (h HttpServer) GetChannelById(w http.ResponseWriter, r *http.Request, channelID string) {}

// Exposes a graphql API endpoint
func (h HttpServer) UseGraphQL(w http.ResponseWriter, r *http.Request) {}

// Opens a WebSocket connection
func (h HttpServer) ConnectWebsocket(w http.ResponseWriter, r *http.Request) {
	// Handles websocket requests from peers.
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Allow connections from any Origin
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("ws: Invalid HTTP method", r.Method)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ws: failed to Upgrade ", err)
		return
	}

	session, _ := h.sessionStore.NewWebsocketSession(conn, &h.service)

	go session.Writer()
	go session.Reader()
}

// Create a user
func (h HttpServer) CreateUser(w http.ResponseWriter, r *http.Request) {}

// Delete a user
func (h HttpServer) DeleteUserById(w http.ResponseWriter, r *http.Request, userID string) {}

// Retrieve a user
func (h HttpServer) GetUserById(w http.ResponseWriter, r *http.Request, userID string) {}
