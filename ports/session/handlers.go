package session

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Muchogoc/semezana/dto"
	"github.com/Muchogoc/semezana/internal/auth"
)

type handler func(ctx context.Context, payload *dto.ClientPayload)

func (s *Session) authMiddleware(next handler) handler {
	return func(ctx context.Context, payload *dto.ClientPayload) {
		token, err := auth.ParseToken(ctx, payload.Auth.Token)
		if err != nil {
			response := &dto.ServerResponse{
				Control: &dto.Ctrl{
					Code:      http.StatusBadRequest,
					Message:   fmt.Errorf("invalid token: %w", err).Error(),
					Timestamp: payload.Timestamp,
				},
			}
			s.queueOut(response)
			return
		}

		err = auth.ValidateToken(ctx, token)
		if err != nil {
			response := &dto.ServerResponse{
				Control: &dto.Ctrl{
					Code:      http.StatusUnauthorized,
					Message:   fmt.Errorf("token validation failed: %w", err).Error(),
					Timestamp: payload.Timestamp,
				},
			}
			s.queueOut(response)
			return
		}

		ctx = auth.SetTokenContext(ctx, token)
		ctx = SetSessionContext(ctx, s)

		next(ctx, payload)
	}
}

func (s *Session) dispatch(ctx context.Context, payload *dto.ClientPayload) {

	var handler handler

	// Add handler variable based on the message type
	switch payload.Type {
	case dto.ClientPayloadTypeHello:
		handler = s.authMiddleware(s.helloHandler)
	case dto.ClientPayloadTypePublish:
		handler = s.authMiddleware(s.publishHandler)
	case dto.ClientPayloadTypeNotify:
		handler = s.notifyHandler
	default:
		log.Println("s.dispatch: unknown message type", s.sid)
		return
	}

	payload.Timestamp = time.Now()

	handler(ctx, payload)
}

func (s *Session) helloHandler(ctx context.Context, payload *dto.ClientPayload) {
	response := s.service.HandleHello(ctx, payload)
	s.queueOut(response)
}

func (s *Session) publishHandler(ctx context.Context, payload *dto.ClientPayload) {
	response := s.service.HandleNewMessage(ctx, payload)
	s.queueOut(response)

}

func (s *Session) notifyHandler(ctx context.Context, payload *dto.ClientPayload) {}
