package session

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Muchogoc/semezana/dto"
)

type handler func(ctx context.Context, payload *dto.ClientPayload)

func (s *Session) dispatch(ctx context.Context, payload *dto.ClientPayload) {

	var handler handler

	// Add handler variable based on the message type
	switch {
	case payload.Hello != nil:
		handler = s.helloHandler
	case payload.Account != nil:
		handler = s.accountHandler
	case payload.Login != nil:
		handler = s.loginHandler
	case payload.Subscription != nil:
		handler = s.subscriptionHandler
	case payload.Leave != nil:
		handler = s.leaveHandler
	case payload.Publish != nil:
		handler = s.publishHandler
	case payload.Delete != nil:
		handler = s.deleteHandler
	case payload.Notify != nil:
		handler = s.notifyHandler

	default:
		log.Println("s.dispatch: unknown message", s.sid)
		return
	}

	payload.Timestamp = time.Now()
	// payload.session = s
	handler(ctx, payload)
}

func (s *Session) helloHandler(ctx context.Context, payload *dto.ClientPayload) {
	response := &dto.ServerResponse{
		Control: &dto.Ctrl{
			Code:      http.StatusOK,
			Timestamp: payload.Timestamp,
		},
	}
	s.queueOut(response)
}

// accountHandler
// 1. Creates a user's account if 'new'
// 1.1 Create a "me" channel on NSQ
// 2. Updates a user's account
func (s *Session) accountHandler(ctx context.Context, payload *dto.ClientPayload) {
	var response *dto.ServerResponse

	s.queueOut(response)
}

// loginHandler
// 1. Get User
// 2. Get user's Subscriptions/Memberships
// 3. For each membership set up a pubsub subscription to it's channel
// i.e subscribe to the channels topic
func (s *Session) loginHandler(ctx context.Context, payload *dto.ClientPayload) {

}

// subscriptionHandler
// 1. Create user subscription to a channel
// 2. if channel doesn't exist create
// 3. Create subscription-channel as NSQ Topic
// 4. Add subscription to session, creates a subsequent channel to subscription-channel
//
// 1. Get User
// 2. Check if channel exists
// 2.1 If channel exists check if there is a membership to the channel
// 2.1.1 GetOrCreate membership
// 2.2 If no channel create channel
// 2.2.1 Create subscription/membership
// 3. Notify other members
// 4. Subscribe to Membership Topic
func (s *Session) subscriptionHandler(ctx context.Context, payload *dto.ClientPayload) {

	response := &dto.ServerResponse{
		Control: &dto.Ctrl{
			Code:    http.StatusOK,
			Message: "",
		},
	}
	s.queueOut(response)
}

// publishHandler
// 1. Get user
// 2. Fetch channel
// 3. Fetch memberships
// 4. Save Message to DB with sequence, author is user
// 5. Save new channel sequence
// 6. Save message audience
// 7. Publish details to membership topics
// 8. Return data response
func (s *Session) publishHandler(ctx context.Context, payload *dto.ClientPayload) {

	response := &dto.ServerResponse{
		Control: &dto.Ctrl{
			Code:    http.StatusOK,
			Message: "",
		},
	}
	s.queueOut(response)

}

func (s *Session) leaveHandler(ctx context.Context, payload *dto.ClientPayload) {}

func (s *Session) deleteHandler(ctx context.Context, payload *dto.ClientPayload) {}

func (s *Session) notifyHandler(ctx context.Context, payload *dto.ClientPayload) {}
