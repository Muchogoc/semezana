package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/dto"
	"github.com/Muchogoc/semezana/ent/schema"
	"github.com/Muchogoc/semezana/internal/auth"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/Muchogoc/semezana/app"

type ChatService struct {
	chatRepo   chat.Repository
	publisher  Publish
	subscriber Subscribe
}

func NewChatService(
	chatRepo chat.Repository,
	publisher Publish,
	subscriber Subscribe,
) ChatService {
	if chatRepo == nil {
		panic("missing chat repository")
	}

	if publisher == nil {
		panic("missing publisher")
	}

	if subscriber == nil {
		panic("missing subscriber")
	}

	service := ChatService{
		chatRepo:   chatRepo,
		publisher:  publisher,
		subscriber: subscriber,
	}

	if err := service.ensureTopicsExist(context.Background()); err != nil {
		panic(fmt.Errorf("failed to create topics: %w", err))
	}

	return service
}

func (c ChatService) ensureTopicsExist(ctx context.Context) error {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "ensureTopicsExist()")
	defer span.End()

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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GenerateAccessToken()")
	defer span.End()

	user, err := c.chatRepo.GetUser(ctx, creds.UserID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	token, err := auth.CreateToken(user.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &dto.TokenResponse{
		Access: token,
	}, nil
}

func (c ChatService) CreateUser(ctx context.Context, newUser dto.NewUser) (*dto.User, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateUser()")
	defer span.End()

	usr := &chat.User{}
	usr.SetID(uuid.NewString())
	usr.SetName(newUser.Name)

	err := c.chatRepo.CreateUser(ctx, usr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &dto.User{
		Id:   usr.ID(),
		Name: usr.Name(),
	}, nil
}

func (c ChatService) GetUsers(ctx context.Context) (*[]dto.User, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUsers()")
	defer span.End()

	users, err := c.chatRepo.GetUsers(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUserById()")
	defer span.End()

	user, err := c.chatRepo.GetUser(ctx, userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &dto.User{
		Id:   user.ID(),
		Name: user.Name(),
	}, nil
}

func (c ChatService) CreateChannel(ctx context.Context, channel dto.NewChannel) (dto.Channel, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateChannel()")
	defer span.End()

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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetChannels()")
	defer span.End()

	channels, err := c.chatRepo.GetChannels(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetChannelById()")
	defer span.End()

	channel, err := c.chatRepo.GetChannel(ctx, channelID, true)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateMembership()")
	defer span.End()

	msp := &chat.Membership{}
	msp.SetID(uuid.NewString())

	user, err := c.chatRepo.GetUser(ctx, *newMembership.UserID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	msp.SetUser(*user)

	channel, err := c.chatRepo.GetChannel(ctx, *newMembership.ChannelID, false)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	msp.SetChannel(*channel)

	err = c.chatRepo.CreateMembership(ctx, msp)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	msg := dto.PubMessage{
		Sender: user.ID(),
		Type:   dto.MessageTypeNewMembership,
		Data:   nil,
	}
	err = c.publisher.PublishToMembership(ctx, msp.ID(), msg)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "HandleHello()")
	defer span.End()

	uid, err := auth.GetUIDFromContext(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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

func (c ChatService) BroadcastMessage(nctx context.Context, message *chat.Message, memberships *[]chat.Membership) {
	// Change from the request context
	ctx := context.Background()
	ctx = trace.ContextWithSpan(ctx, trace.SpanFromContext(nctx))

	ctx, span := otel.Tracer(tracerName).Start(ctx, "BroadcastMessage()")
	defer span.End()

	author := message.Author()
	channel := message.Channel()

	send := func(ctx context.Context, m chat.Membership) {
		membership, err := c.chatRepo.GetMembership(ctx, m.ID(), true)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Printf("failed to get membership: %s", err.Error())
			return
		}

		user := membership.User()

		recipient := chat.Recipient{}
		recipient.SetUserID(user.ID())
		recipient.SetMessageID(message.ID())
		recipient.SetStatus("DELIVERED")
		recipient.SetStatusAt(time.Now())

		err = c.chatRepo.CreateRecipient(ctx, &recipient)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Printf("failed to create recipient: %s", err.Error())
			return
		}

		input := dto.PubMessage{
			Sender: author.ID(),
			Type:   dto.MessageTypeNewMessage,
			Data: dto.Data{
				Head:      schema.MessageHeaders{},
				Channel:   channel.ID(),
				From:      author.ID(),
				Timestamp: message.Timestamp(),
				Sequence:  1,
				Content: schema.MessageContent{
					Text: message.Content().Text(),
				},
			},
		}

		err = c.publisher.PublishToMembership(ctx, membership.ID(), input)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Printf("failed to publish to membership: %s", err.Error())
			return
		}
	}

	for _, m := range *memberships {
		go send(ctx, m)
	}

}

func (c ChatService) HandleNewMessage(ctx context.Context, payload *dto.ClientPayload) *dto.ServerResponse {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "HandleNewMessage()")
	defer span.End()

	uid, err := auth.GetUIDFromContext(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	message, memberships, err := channel.NewMessage(payload.Publish.Content.(string), uid, payload.Timestamp)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return &dto.ServerResponse{
			Control: &dto.Ctrl{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				Timestamp: payload.Timestamp,
			},
		}
	}

	go c.BroadcastMessage(ctx, message, memberships)

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
	_, span := otel.Tracer(tracerName).Start(ctx, "ProcessPubsubMessage()")
	defer span.End()

	switch payload.Type {
	case dto.MessageTypeNewMessage:
		data := &dto.Data{}
		_ = mapstructure.Decode(payload.Data, data)

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
