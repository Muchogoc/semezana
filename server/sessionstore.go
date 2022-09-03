package main

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// NewSessionStore initializes a session store.
func NewSessionStore() *SessionStore {
	store := &SessionStore{
		sessions: make(map[string]*Session),
	}

	return store
}

// SessionStore holds live sessions. Sessions are stored in a map indexed by session ID.
type SessionStore struct {
	lock sync.Mutex
	// All sessions indexed by session ID
	sessions map[string]*Session
}

// NewSession creates a new session and saves it to the session store.
func (s *SessionStore) NewSession(conn *websocket.Conn) (*Session, int) {
	var session Session

	session.sid = uuid.NewString()

	s.lock.Lock()
	if _, found := s.sessions[session.sid]; found {
		log.Fatalln("ERROR! duplicate session ID", session.sid)
	}
	s.lock.Unlock()

	session.ws = conn
	session.subscriptions = make(map[string]*Subscription)
	session.send = make(chan interface{}, sendQueueLimit+32) // buffered
	session.stop = make(chan interface{}, 1)                 // Buffered by 1 just to make it non-blocking
	session.detach = make(chan string, 64)                   // buffered

	s.lock.Lock()

	s.sessions[session.sid] = &session
	numSessions := len(s.sessions)

	s.lock.Unlock()

	return &session, numSessions
}

// Get fetches a session from store by session ID.
func (s *SessionStore) Get(id string) *Session {
	s.lock.Lock()
	defer s.lock.Unlock()

	if session := s.sessions[id]; session != nil {
		return session
	}

	return nil
}

// Delete removes session from store.
func (s *SessionStore) Delete(session *Session) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.sessions, session.sid)
}

// Shutdown terminates sessionStore.
func (s *SessionStore) Shutdown() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, session := range s.sessions {
		session.stopSession(1)
	}
}
