package semezana

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/subscription"
	"github.com/Muchogoc/semezana/ent/topic"
	"github.com/Muchogoc/semezana/semezana/dto"
	"github.com/Muchogoc/semezana/semezana/models"
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

// accountHandler
// 1. Creates a user's account if 'new'
// 1.1 Create a "me" topic on NSQ
// 2. Updates a user's account
func (s *Session) accountHandler(ctx context.Context, msg *dto.ClientComMessage) {
	var response *dto.ServerComMessage

	if strings.HasPrefix(msg.Account.User, "new") {
		user, err := globals.client.User.
			Create().
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

// loginHandler
// 1. Set up session
// 2. Register session as channel to user's subscription-topics
func (s *Session) loginHandler(ctx context.Context, msg *dto.ClientComMessage) {
	uid, err := uuid.Parse(msg.Login.User)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	user, err := globals.client.User.Get(ctx, uid)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	subscriptions, err := user.QuerySubscriptions().All(ctx)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	for _, subscription := range subscriptions {
		s.addSub(subscription.TopicID.String(), subscription)
		go s.nsqConsumer(ctx, subscription)
	}

}

// subscriptionHandler
// 1. Create user subscription to a topic
// 2. if topic doesn't exist create
// 3. Create subscription-topic as NSQ Topic
// 4. Add subscription to session, creates a subsequent channel to subscription-topic
func (s *Session) subscriptionHandler(ctx context.Context, msg *dto.ClientComMessage) {
	uid, err := uuid.Parse(msg.Subscription.User)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	user, err := globals.client.User.Get(ctx, uid)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	cid, err := uuid.Parse(msg.Subscription.Topic)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	// check if topic exists
	exists, err := globals.client.Topic.Query().Where(
		topic.ID(cid),
	).Exist(ctx)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	var sub *ent.Subscription
	if exists {
		// check if there is a subscription with the topic
		exist, err := user.QuerySubscriptions().Where(
			subscription.TopicID(cid),
		).Exist(ctx)
		if err != nil {
			response := &dto.ServerComMessage{
				Control: &dto.MsgServerCtrl{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}
			s.queueOut(response)
			return
		}

		if !exist {
			sub, err = globals.client.Subscription.Create().
				SetSubscriber(user).
				SetTopicID(cid).
				Save(ctx)
			if err != nil {
				response := &dto.ServerComMessage{
					Control: &dto.MsgServerCtrl{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					},
				}
				s.queueOut(response)
				return
			}
		} else {
			sub, err = user.QuerySubscriptions().Where(
				subscription.TopicID(cid),
			).Only(ctx)
			if err != nil {
				response := &dto.ServerComMessage{
					Control: &dto.MsgServerCtrl{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					},
				}
				s.queueOut(response)
				return
			}
		}

	} else {
		// create topic
		topic, err := globals.client.Topic.Create().
			SetName("Test Topic").
			SetType(string(models.TopicCategoryP2P)).
			SetState(models.StateOK.String()).
			SetStateAt(time.Now()).
			Save(ctx)
		if err != nil {
			response := &dto.ServerComMessage{
				Control: &dto.MsgServerCtrl{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}
			s.queueOut(response)
			return
		}

		// create subscription
		// add subscription to user
		sub, err = globals.client.Subscription.Create().
			SetSubscriber(user).
			SetTopic(topic).
			Save(ctx)
		if err != nil {
			response := &dto.ServerComMessage{
				Control: &dto.MsgServerCtrl{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}
			s.queueOut(response)
			return
		}
	}

	m := PubMessage{
		Type: "subscription.create",
	}

	err = globals.producer.Publish(sub.ID.String(), m.Marshal())
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	s.addSub(sub.TopicID.String(), sub)
	go s.nsqConsumer(ctx, sub)

	response := &dto.ServerComMessage{
		Control: &dto.MsgServerCtrl{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("topic: %s, subscription: %s", sub.TopicID.String(), sub.ID.String()),
		},
	}
	s.queueOut(response)
}

// publishHandler
// 1. Fetch topic
// 2. Fetch subscription-topic
// 2. Save Message to DB
// 3. Publish details to topic
// 4. Return data response
func (s *Session) publishHandler(ctx context.Context, msg *dto.ClientComMessage) {
	uid, err := uuid.Parse(msg.Publish.User)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	cid, err := uuid.Parse(msg.Publish.Topic)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	topic, err := globals.client.Topic.Get(ctx, cid)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	sequence := topic.Sequence + 1

	message, err := globals.client.Message.Create().
		SetSenderID(uid).
		SetTopic(topic).
		SetHeader(msg.Publish.Head).
		SetContent(map[string]interface{}{
			"content": msg.Publish.Content,
		}).
		SetSequence(sequence).
		Save(ctx)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	topic, err = topic.Update().
		SetSequence(sequence).
		SetTouched(msg.Timestamp).
		Save(ctx)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	subscriptions, err := topic.QuerySubscriptions().All(ctx)
	if err != nil {
		response := &dto.ServerComMessage{
			Control: &dto.MsgServerCtrl{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
		s.queueOut(response)
		return
	}

	payload := PubMessage{
		Type:    "message.new",
		Sender:  uid.String(),
		Message: message,
	}

	for _, subscription := range subscriptions {
		globals.producer.Publish(subscription.ID.String(), payload.Marshal())
	}

}

func (s *Session) leaveHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) deleteHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) notifyHandler(ctx context.Context, msg *dto.ClientComMessage) {}
