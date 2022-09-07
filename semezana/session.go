package semezana

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/Muchogoc/semezana/semezana/dto"
	"github.com/gorilla/websocket"
)

// Maximum number of queued messages before session is considered stale and dropped.
const sendQueueLimit = 128

// Session holds and represents a single WS connection to the server
// It is responsible for maintaining the connection to a client
type Session struct {
	// ID of the current session
	sid string
	ws  *websocket.Conn

	// Channel for shutting down the session, buffer 1.
	// Content in the same format as for 'send'
	stop chan interface{}
	// Outbound messages channel, buffered.
	send chan interface{}

	// messageLock *sync.Mutex

	// detach - channel for detaching session from channel, buffered.
	// Content is channel name to detach from.
	detach chan string

	// ID of the current user. Could be empty if session is not authenticated
	uid string
	// Map of channel subscriptions, indexed by channel name.
	// Don't access directly. Use getters/setters.
	// subscriptions map[string]*Subscription
	// subsLock      sync.RWMutex
	subscriptions *sync.Map

	// connected client details
	// IP address of the.
	remoteAddress string
}

// cleanUp is called when the session is terminated to perform resource cleanup.
func (s *Session) cleanUp() {
	s.stop <- 1
}

func (s *Session) stopSession(data interface{}) {
	s.stop <- data
}

// Message received, convert bytes to ClientComMessage and dispatch
func (s *Session) dispatchRaw(ctx context.Context, raw []byte) {
	var msg dto.ClientComMessage

	if err := json.Unmarshal(raw, &msg); err != nil {
		log.Println("s.dispatch", err, s.sid)
	}

	s.dispatch(ctx, &msg)
}

// queueOut attempts to send a ServerComMessage to a session write loop;
// it fails, if the send buffer is full.
func (s *Session) queueOut(msg *dto.ServerComMessage) bool {
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
