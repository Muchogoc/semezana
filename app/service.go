package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/domain/user"
	"github.com/Muchogoc/semezana/dto"
	"github.com/Muchogoc/semezana/ent/schema"
	"github.com/Muchogoc/semezana/internal/auth"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type ChatService struct {
	chatRepo   chat.Repository
	userRepo   user.Repository
	publisher  Publish
	subscriber Subscribe
}

func NewChatService(
	chatRepo chat.Repository,
	userRepo user.Repository,
	publisher Publish,
	subscriber Subscribe,
) ChatService {
	if chatRepo == nil {
		panic("missing chat repository")
	}

	if userRepo == nil {
		panic("missing user repository")
	}

	if publisher == nil {
		panic("missing publisher")
	}

	if subscriber == nil {
		panic("missing subscriber")
	}

	service := ChatService{
		chatRepo:   chatRepo,
		userRepo:   userRepo,
		publisher:  publisher,
		subscriber: subscriber,
	}

	if err := service.ensureTopicsExist(context.Background()); err != nil {
		panic(fmt.Errorf("failed to create topics: %w", err))
	}

	return service
}

func (c ChatService) ensureTopicsExist(ctx context.Context) error {
	memberships, err := c.chatRepo.GetAllMemberships(ctx)
	if err != nil {
		return err
	}

	for _, membership := range memberships {
		msg := dto.PubMessage{
			Sender: "SYSTEM",
			Type:   dto.MessageTypeNewMembership,
			Data:   nil,
		}
		err = c.publisher.PublishToMembership(ctx, membership, msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c ChatService) GenerateAccessToken(ctx context.Context, creds dto.NewToken) (*dto.TokenResponse, error) {
	user, err := c.userRepo.GetUser(ctx, creds.UserID)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(user.ID())
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		Access: token,
	}, nil
}

func (c ChatService) CreateUser(ctx context.Context, newUser dto.NewUser) (*dto.User, error) {

	usr := &user.User{}
	usr.SetID(uuid.NewString())
	usr.SetName(newUser.Name)

	err := c.userRepo.CreateUser(ctx, usr)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		Id:   usr.ID(),
		Name: usr.Name(),
	}, nil
}

func (c ChatService) GetUsers(ctx context.Context) (*[]dto.User, error) {
	users, err := c.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usrs []dto.User
	for _, user := range *users {
		usr := dto.User{
			Id:   user.ID(),
			Name: user.Name(),
		}

		usrs = append(usrs, usr)
	}

	return &usrs, nil
}

func (c ChatService) GetUserById(ctx context.Context, userID string) (*dto.User, error) {
	user, err := c.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		Id:   user.ID(),
		Name: user.Name(),
	}, nil
}

func (c ChatService) CreateChannel(ctx context.Context, channel dto.NewChannel) (dto.Channel, error) {
	chn := &chat.Channel{}
	chn.SetID(uuid.NewString())
	chn.SetName(channel.Name)
	if channel.Description != nil {
		chn.SetDescription(*channel.Description)
	}
	chn.SetCategory(chat.ChannelCategory(channel.Category))
	chn.SetState(chat.StateOK)

	err := c.chatRepo.CreateChannel(ctx, chn)
	if err != nil {
		return dto.Channel{}, err
	}

	description := chn.Description()
	return dto.Channel{
		Id:          chn.ID(),
		Name:        chn.Name(),
		Category:    dto.ChannelCategory(chn.Category()),
		Description: &description,
	}, nil
}

func (c ChatService) GetChannels(ctx context.Context) (*[]dto.Channel, error) {
	channels, err := c.chatRepo.GetChannels(ctx)
	if err != nil {
		return nil, err
	}

	var chns []dto.Channel
	for _, chn := range *channels {
		description := chn.Description()
		usr := dto.Channel{
			Id:          chn.ID(),
			Name:        chn.Name(),
			Category:    dto.ChannelCategory(chn.Category()),
			Description: &description,
		}

		chns = append(chns, usr)
	}

	return &chns, nil
}

func (c ChatService) GetChannelById(ctx context.Context, channelID string) (*dto.Channel, error) {
	channel, err := c.chatRepo.GetChannel(ctx, channelID, true)
	if err != nil {
		return nil, err
	}

	description := channel.Description()
	return &dto.Channel{
		Id:          channel.ID(),
		Name:        channel.Name(),
		Category:    dto.ChannelCategory(channel.Category()),
		Description: &description,
	}, nil
}

func (c ChatService) CreateMembership(ctx context.Context, newMembership dto.NewMembership) (*dto.Membership, error) {

	msp := &chat.Membership{}
	msp.SetID(uuid.NewString())

	user, err := c.userRepo.GetUser(ctx, *newMembership.UserID)
	if err != nil {
		return nil, err
	}
	msp.SetUser(*user)

	channel, err := c.chatRepo.GetChannel(ctx, *newMembership.ChannelID, false)
	if err != nil {
		return nil, err
	}
	msp.SetChannel(*channel)

	err = c.chatRepo.CreateMembership(ctx, msp)
	if err != nil {
		return nil, err
	}

	msg := dto.PubMessage{
		Sender: user.ID(),
		Type:   dto.MessageTypeNewMembership,
		Data:   nil,
	}
	err = c.publisher.PublishToMembership(ctx, msp.ID(), msg)
	if err != nil {
		return nil, err
	}

	cid := channel.ID()
	uid := user.ID()
	return &dto.Membership{
		Id:        msp.ID(),
		ChannelID: &cid,
		UserID:    &uid,
	}, nil
}

// get memberships
// getOrCreate device
// set current device
// set up pubsub subscribers
func (c ChatService) HandleHello(ctx context.Context, payload *dto.ClientPayload) *dto.ServerResponse {
	uid, err := auth.GetUIDFromContext(ctx)
	if err != nil {
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	memberships, err := c.chatRepo.GetUserMemberships(ctx, uid)
	if err != nil {
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	for _, membership := range *memberships {
		go c.subscriber.CreateSessionSubscriber(ctx, membership.ID())
	}

	return &dto.ServerResponse{
		Control: &dto.Ctrl{
			Code:      http.StatusOK,
			Timestamp: payload.Timestamp,
		},
	}
}

func (c ChatService) HandleNewMessage(ctx context.Context, payload *dto.ClientPayload) *dto.ServerResponse {
	uid, err := auth.GetUIDFromContext(ctx)
	if err != nil {
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	channel, err := c.chatRepo.GetChannel(ctx, payload.Publish.Channel, true)
	if err != nil {
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	message, memberships, err := channel.NewMessage(payload.Publish.Content.(string), uid)
	if err != nil {
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	err = c.chatRepo.CreateMessage(ctx, message, uid, channel.ID())
	if err != nil {
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	for _, m := range *memberships {
		membership, err := c.chatRepo.GetMembership(ctx, m.ID(), true)
		if err != nil {
			return &dto.ServerResponse{
				Control: &dto.Ctrl{
					Code:      http.StatusInternalServerError,
					Message:   err.Error(),
					Timestamp: payload.Timestamp,
				},
			}
		}

		user := membership.User()

		recipient := chat.Recipient{}
		recipient.SetUserID(user.ID())
		recipient.SetMessageID(message.ID())
		recipient.SetStatus("DELIVERED")
		recipient.SetStatusAt(time.Now())

		err = c.chatRepo.CreateRecipient(ctx, &recipient)
		if err != nil {
			return &dto.ServerResponse{
				Control: &dto.Ctrl{
					Code:      http.StatusInternalServerError,
					Message:   err.Error(),
					Timestamp: payload.Timestamp,
				},
			}
		}

		input := dto.PubMessage{
			Sender: uid,
			Type:   dto.MessageTypeNewMessage,
			Data: dto.Data{
				Head:      schema.MessageHeaders{},
				Channel:   channel.ID(),
				From:      uid,
				Timestamp: payload.Timestamp,
				Sequence:  1,
				Content: schema.MessageContent{
					Text: message.Content().Text(),
				},
			},
		}
		err = c.publisher.PublishToMembership(ctx, membership.ID(), input)
		if err != nil {
			return &dto.ServerResponse{
				Control: &dto.Ctrl{
					Code:      http.StatusInternalServerError,
					Message:   err.Error(),
					Timestamp: payload.Timestamp,
				},
			}
		}

	}

	return &dto.ServerResponse{
		Control: &dto.Ctrl{
			Code:      http.StatusOK,
			Timestamp: payload.Timestamp,
		},
		Data: &dto.Data{
			Head:      schema.MessageHeaders{},
			Channel:   channel.ID(),
			From:      uid,
			Timestamp: payload.Timestamp,
			Sequence:  1,
			Content: schema.MessageContent{
				Text: message.Content().Text(),
			},
		},
	}
}

func (c ChatService) ProcessPubsubMessage(ctx context.Context, payload dto.PubMessage) *dto.ServerResponse {
	switch payload.Type {
	case dto.MessageTypeNewMessage:
		data := &dto.Data{}
		mapstructure.Decode(payload.Data, data)

		response := &dto.ServerResponse{
			Type: dto.ServerResponseTypeData,
			Data: data,
			Control: &dto.Ctrl{
				Code: http.StatusOK,
			},
		}

		return response

	case dto.MessageTypeNewMembership:
		return nil
	}

	return nil
}
