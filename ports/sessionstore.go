package ports

import (
	"log"
	"sync"

	"github.com/Muchogoc/semezana/app"
	"github.com/Muchogoc/semezana/ports/session"
	"github.com/gorilla/websocket"
)

// NewSessionStore initializes a session store.
func NewSessionStore() *SessionStore {
	store := &SessionStore{
		sessions: make(map[string]*session.Session),
	}

	return store
}

// SessionStore holds live sessions. Sessions are stored in a map indexed by session ID.
type SessionStore struct {
	lock sync.Mutex
	// All sessions indexed by session ID
	sessions map[string]*session.Session
}

// NewSession creates a new session and saves it to the session store.
func (s *SessionStore) NewWebsocketSession(conn *websocket.Conn, app *app.ChatService) (*session.Session, int) {
	session := session.NewWebsocketSession(conn, app)

	s.lock.Lock()

	if _, found := s.sessions[session.ID()]; found {
		log.Fatalln("ERROR! duplicate session ID", session.ID())
	}

	s.sessions[session.ID()] = session
	numSessions := len(s.sessions)

	s.lock.Unlock()

	return session, numSessions
}

// Get fetches a session from store by session ID.
func (s *SessionStore) Get(id string) *session.Session {
	s.lock.Lock()
	defer s.lock.Unlock()

	if session := s.sessions[id]; session != nil {
		return session
	}

	return nil
}

// Delete removes session from store.
func (s *SessionStore) Delete(session *session.Session) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.sessions, session.ID())
}

// Shutdown terminates sessionStore.
func (s *SessionStore) Shutdown() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, session := range s.sessions {
		session.StopSession(1)
	}
}
