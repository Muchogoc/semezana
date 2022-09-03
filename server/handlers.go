package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Muchogoc/semezana/server/dto"
	"github.com/Muchogoc/semezana/server/models"
	"github.com/google/uuid"
)

type handler func(ctx context.Context, msg *dto.ClientComMessage)

func (s *Session) dispatch(ctx context.Context, msg *dto.ClientComMessage) {

	var handler handler

	// Add handler variable based on the message type
	switch {
	case msg.Hi != nil:
		handler = s.handshakeHandler
	case msg.Account != nil:
		handler = s.accountHandler
	case msg.Login != nil:
		handler = s.loginHandler
	case msg.Subscription != nil:
		handler = s.subscriptionHandler
	case msg.Leave != nil:
		handler = s.leaveHandler
	case msg.Publish != nil:
		handler = s.publishHandler
	case msg.Delete != nil:
		handler = s.deleteHandler
	case msg.Notify != nil:
		handler = s.notifyHandler

	default:
		log.Println("s.dispatch: unknown message", s.sid)
		return
	}

	msg.Timestamp = time.Now()
	// msg.session = s
	handler(ctx, msg)
}

func (s *Session) handshakeHandler(ctx context.Context, msg *dto.ClientComMessage) {
	output := &dto.ServerComMessage{
		Control: &dto.MsgServerCtrl{
			Code:      http.StatusOK,
			Timestamp: msg.Timestamp,
		},
	}
	s.queueOut(output)
}

func (s *Session) accountHandler(ctx context.Context, msg *dto.ClientComMessage) {
	var response *dto.ServerComMessage

	if strings.HasPrefix(msg.Account.User, "new") {
		user, err := globals.client.User.
			Create().
			SetID(uuid.New()).
			SetState(models.StateOK.String()).
			SetStateAt(msg.Timestamp).
			Save(ctx)
		if err != nil {
			response = &dto.ServerComMessage{
				Control: &dto.MsgServerCtrl{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}
			s.queueOut(response)
			return
		}

		g, _ := json.Marshal(user)
		log.Println("user:", string(g))

		response = &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:      http.StatusCreated,
				Timestamp: msg.Timestamp,
				Parameters: map[string]interface{}{
					"user":      user.ID,
					"createdAt": user.CreatedAt,
					"updatedAt": user.UpdatedAt,
				},
			},
		}

	}

	s.queueOut(response)
}

func (s *Session) loginHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) subscriptionHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) leaveHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) publishHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) deleteHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) notifyHandler(ctx context.Context, msg *dto.ClientComMessage) {}
