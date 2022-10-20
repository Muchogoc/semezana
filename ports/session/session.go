package session

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/Muchogoc/semezana/app"
	"github.com/Muchogoc/semezana/dto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Maximum number of queued messages before session is considered stale and dropped.
const sendQueueLimit = 128

// Session holds and represents a single connection to the server
// It is responsible for maintaining the connection to a client
type Session struct {
	// ID of the session
	sid string
	ws  *websocket.Conn

	// Channel for shutting down the session, buffer 1.
	// Content in the same format as for 'send'
	stop chan interface{}
	// Outbound messages channel, buffered.
	send chan interface{}

	// Inbound messages from a subscriber
	sub chan dto.PubMessage

	// messageLock *sync.Mutex

	// detach - channel for detaching session from channel, buffered.
	// Content is channel name to detach from.
	detach chan string

	// ID of the current user. Could be empty if session is not authenticated
	// uid string

	// Map of channel subscriptions/memberships, indexed by channel name.
	// Don't access directly. Use getters/setters.
	subscriptions *sync.Map
	service       *app.ChatService
}

func (s Session) ID() string {
	return s.sid
}

func (s Session) StopChan() chan interface{} {
	return s.stop
}

// cleanUp is called when the session is terminated to perform resource cleanup.
func (s *Session) cleanUp() {
	s.stop <- 1
}

func (s *Session) StopSession(data interface{}) {
	s.stop <- data
}

// Message received, convert bytes to ClientPayload and dispatch
func (s *Session) dispatchRaw(ctx context.Context, raw []byte) {
	var msg dto.ClientPayload

	if err := json.Unmarshal(raw, &msg); err != nil {
		log.Println("s.dispatch", err, s.sid)
	}

	s.dispatch(ctx, &msg)
}

// queueOut attempts to send a ServerResponse to a session write loop;
// it fails, if the send buffer is full.
func (s *Session) queueOut(msg *dto.ServerResponse) bool {
	if s == nil {
		return true
	}

	select {
	case s.send <- msg:
	default:
		// Never block here since it may also block the channel's run() goroutine.
		return false
	}

	return true
}

func NewWebsocketSession(conn *websocket.Conn, service *app.ChatService) *Session {
	return &Session{
		sid:           uuid.NewString(),
		ws:            conn,
		stop:          make(chan interface{}, 1),
		send:          make(chan interface{}, sendQueueLimit+32),
		detach:        make(chan string, 64),
		subscriptions: &sync.Map{},
		service:       service,
	}
}
