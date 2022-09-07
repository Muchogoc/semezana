package semezana

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/channel"
	"github.com/Muchogoc/semezana/ent/subscription"
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
// 1.1 Create a "me" channel on NSQ
// 2. Updates a user's account
func (s *Session) accountHandler(ctx context.Context, msg *dto.ClientComMessage) {
	var response *dto.ServerComMessage

	if strings.HasPrefix(msg.Account.User, "new") {
		user, err := globals.client.User.
			Create().
			SetName("Test User").
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
// 2. Register session as channel to user's subscription-channels
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
		s.addSub(subscription.ChannelID.String(), subscription)
		go s.nsqConsumer(ctx, subscription)
	}

}

// subscriptionHandler
// 1. Create user subscription to a channel
// 2. if channel doesn't exist create
// 3. Create subscription-channel as NSQ Topic
// 4. Add subscription to session, creates a subsequent channel to subscription-channel
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

	cid, err := uuid.Parse(msg.Subscription.Channel)
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

	// check if channel exists
	exists, err := globals.client.Channel.Query().Where(channel.ID(cid)).Exist(ctx)
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
		// check if there is a subscription with the channel
		exist, err := user.QuerySubscriptions().Where(
			subscription.ChannelID(cid),
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
				SetRole("OWNER").
				SetStatus("OK").
				SetUser(user).
				SetChannelID(cid).
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
				subscription.ChannelID(cid),
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
		// create channel
		channel, err := globals.client.Channel.Create().
			SetName("Test Topic").
			SetType(string(models.ChannelCategoryP2P)).
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
			SetUser(user).
			SetChannel(channel).
			SetRole("OWNER").
			SetStatus("OK").
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

	s.addSub(sub.ChannelID.String(), sub)
	go s.nsqConsumer(ctx, sub)

	response := &dto.ServerComMessage{
		Control: &dto.MsgServerCtrl{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("channel: %s, subscription: %s", sub.ChannelID.String(), sub.ID.String()),
		},
	}
	s.queueOut(response)
}

// publishHandler
// 1. Fetch channel
// 2. Fetch subscription-channel
// 2. Save Message to DB
// 3. Publish details to channel
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

	cid, err := uuid.Parse(msg.Publish.Channel)
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

	channel, err := globals.client.Channel.Get(ctx, cid)
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

	sequence := channel.Sequence + 1

	message, err := globals.client.Message.Create().
		SetAuthorID(uid).
		SetChannel(channel).
		SetHeader(msg.Publish.Head).
		SetContent(models.MessageContent{
			Text: msg.Publish.Content.(string),
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

	channel, err = channel.Update().
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

	subscriptions, err := channel.QuerySubscriptions().All(ctx)
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
		_, err := globals.client.Recipient.Create().
			SetMessage(message).
			SetUserID(subscription.UserID).
			SetStatus("DELIVERED").
			SetDeliveredAt(msg.Timestamp).
			Save(ctx)
		if err != nil {
			response := &dto.ServerComMessage{
				Control: &dto.MsgServerCtrl{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			}
			s.queueOut(response)
		}

		globals.producer.Publish(subscription.ID.String(), payload.Marshal())
	}

}

func (s *Session) leaveHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) deleteHandler(ctx context.Context, msg *dto.ClientComMessage) {}

func (s *Session) notifyHandler(ctx context.Context, msg *dto.ClientComMessage) {}
