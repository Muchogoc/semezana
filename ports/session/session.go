package session

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Muchogoc/semezana/dto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Maximum number of queued messages before session is considered stale and dropped.
const sendQueueLimit = 128

type Service interface {
	HandleHello(ctx context.Context, payload *dto.ClientPayload) *dto.ServerResponse
	HandleNewMessage(ctx context.Context, payload *dto.ClientPayload) *dto.ServerResponse
	ProcessPubsubMessage(ctx context.Context, payload dto.PubMessage) *dto.ServerResponse
}

// Session holds and represents a single connection to the server
// It is responsible for maintaining the connection to a client
type Session struct {
	// ID of the session
	sid string
	// ID of the session's user.
	uid string

	ws *websocket.Conn

	// Channel for shutting down the session, buffer 1.
	stop chan interface{}

	// Outbound messages channel
	send chan interface{}

	// Inbound messages channel from a pull pubsub subscriber
	// sub chan dto.PubMessage

	ctx context.Context

	service Service
}

func (s Session) ID() string {
	return s.sid
}

func (s Session) User() string {
	return s.uid
}

func (s Session) StopChan() chan interface{} {
	return s.stop
}

// cleanUp is called when the session is terminated to perform resource cleanup.
func (s *Session) cleanUp() {
	s.stop <- 1
}

func (s *Session) Stop(data interface{}) {
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
	if msg == nil {
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

func NewWebsocketSession(conn *websocket.Conn, service Service) *Session {
	return &Session{
		sid:     uuid.NewString(),
		ws:      conn,
		stop:    make(chan interface{}, 1),
		send:    make(chan interface{}, sendQueueLimit+32),
		service: service,
	}
}

func SetSessionContext(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, dto.ContextKeySession, session)
}
