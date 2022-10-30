package session

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/Muchogoc/semezana/dto"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = time.Second * 55

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func (s *Session) closeWebsocket() {
	s.ws.Close()
}

func (s *Session) Reader() {
	defer func() {
		s.closeWebsocket()
		s.cleanUp()
	}()

	// s.ws.SetReadLimit(globals.maxMessageSize)
	s.ws.SetReadDeadline(time.Now().Add(pongWait))
	s.ws.SetPongHandler(
		func(string) error {
			s.ws.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		},
	)

	for {
		_, raw, err := s.ws.ReadMessage()
		if err != nil {
			return
		}
		s.dispatchRaw(context.Background(), raw)
	}
}

func (s *Session) Writer() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		s.closeWebsocket()
	}()

	for {
		select {
		case msg, ok := <-s.send:
			if !ok {
				return
			}

			switch v := msg.(type) {
			case []*dto.ServerResponse:
				for _, msg := range v {
					if err := wsWrite(s.ws, websocket.TextMessage, msg); err != nil {
						return
					}
				}
			case *dto.ServerResponse:
				if err := wsWrite(s.ws, websocket.TextMessage, v); err != nil {
					return
				}
			default:
				return
			}

		case <-s.stop:
			return

		case <-ticker.C:
			if err := wsWrite(s.ws, websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Writes a message with the given message type (mt) and payload.
func wsWrite(ws *websocket.Conn, messageType int, msg interface{}) error {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return err
	}

	ws.SetWriteDeadline(time.Now().Add(writeWait))

	return ws.WriteMessage(messageType, buf.Bytes())
}
