package adapters

import (
	"context"
	"time"

	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/domain/user"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/schema"
	"github.com/Muchogoc/semezana/ent/subscription"
	"github.com/google/uuid"
)

type EntRepository struct {
	client      *ent.Client
	chatFactory chat.Factory
	userFactory user.Factory
}

func NewEntRepository(client *ent.Client, chatFactory chat.Factory, userFactory user.Factory) *EntRepository {
	return &EntRepository{
		client:      client,
		chatFactory: chatFactory,
		userFactory: userFactory,
	}
}

func (e EntRepository) CreateChannel(ctx context.Context, channel *chat.Channel) error {
	cid, err := uuid.Parse(channel.ID())
	if err != nil {
		return err
	}

	nch, err := e.client.Channel.Create().
		SetID(cid).
		SetName(channel.Name()).
		SetDescription(channel.Description()).
		SetType(channel.Category().String()).
		SetState(channel.State().String()).
		SetStateAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	newChannel, err := e.chatFactory.UnmarshalChannelFromDatabase(
		nch.ID.String(),
		nch.Description,
		nch.Name,
		chat.ChannelState(nch.State),
		chat.ChannelCategory(nch.Type),
	)
	if err != nil {
		return err
	}

	channel = newChannel

	return nil

}

func (e EntRepository) GetChannel(ctx context.Context, id string) (*chat.Channel, error) {
	cid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	channel, err := e.client.Channel.Get(ctx, cid)
	if err != nil {
		return nil, err
	}

	return e.chatFactory.UnmarshalChannelFromDatabase(
		channel.ID.String(),
		channel.Description,
		channel.Name,
		chat.ChannelState(channel.State),
		chat.ChannelCategory(channel.Type),
	)
}

func (e EntRepository) GetChannels(ctx context.Context) (*[]chat.Channel, error) {
	chnls, err := e.client.Channel.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var channels []chat.Channel

	for _, chnl := range chnls {
		channel, err := e.chatFactory.UnmarshalChannelFromDatabase(
			chnl.ID.String(),
			chnl.Description,
			chnl.Name,
			chat.ChannelState(chnl.State),
			chat.ChannelCategory(chnl.Type),
		)
		if err != nil {
			return nil, err
		}

		channels = append(channels, *channel)
	}

	return &channels, nil
}

func (e EntRepository) UpdateChannel(
	ctx context.Context,
	id string,
	updateFn func(h *chat.Channel) (*chat.Channel, error),
) error {

	return nil
}

func (e EntRepository) CreateMembership(ctx context.Context, membership *chat.Membership) error {
	user := membership.User()
	uid, err := uuid.Parse(user.ID())
	if err != nil {
		return err
	}

	channel := membership.Channel()
	cid, err := uuid.Parse(channel.ID())
	if err != nil {
		return err
	}

	_, err = e.client.Subscription.Create().
		SetUserID(uid).
		SetChannelID(cid).
		SetRole(membership.Role().String()).
		SetStatus("OK").
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e EntRepository) GetMembership(ctx context.Context, userID, channelID string) (*chat.Membership, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	cid, err := uuid.Parse(channelID)
	if err != nil {
		return nil, err
	}

	sub, err := e.client.Subscription.Query().Where(
		subscription.UserID(uid),
		subscription.ChannelID(cid),
	).Only(ctx)
	if err != nil {
		return nil, err
	}

	return e.chatFactory.UnmarshalMembershipFromDatabase(
		sub.ID.String(),
		chat.MembershipRole(sub.Role),
	)
}

func (e EntRepository) UpdateMembership(
	ctx context.Context,
	id string,
	updateFn func(h *chat.Membership) (*chat.Membership, error),
) error {
	return nil
}

func (e EntRepository) CreateMessage(ctx context.Context, message *chat.Message, userID, channelID string) error {
	author := message.Author()
	uid, err := uuid.Parse(author.ID())
	if err != nil {
		return err
	}

	channel := message.Channel()
	cid, err := uuid.Parse(channel.ID())
	if err != nil {
		return err
	}

	content := message.Content()

	_, err = e.client.Message.Create().
		SetAuthorID(uid).
		SetChannelID(cid).
		SetHeader(schema.MessageHeaders{}).
		SetContent(schema.MessageContent{
			Text: content.Text(),
		}).
		SetSequence(1).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (e EntRepository) CreateMessageAudience(ctx context.Context, audience *chat.Audience) error {
	return nil
}

func (e EntRepository) CreateUser(ctx context.Context, user *user.User) error {
	uid, err := uuid.Parse(user.ID())
	if err != nil {
		return err
	}

	usr, err := e.client.User.
		Create().
		SetID(uid).
		SetName(user.Name()).
		SetState("OK").
		SetStateAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	new, err := e.userFactory.UnmarshalUserFromDatabase(
		usr.ID.String(), usr.Name,
	)
	if err != nil {
		return err
	}

	user = new

	return nil
}

func (e EntRepository) GetUser(ctx context.Context, id string) (*user.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	usr, err := e.client.User.Get(ctx, uid)
	if err != nil {
		return nil, err
	}

	return e.userFactory.UnmarshalUserFromDatabase(
		usr.ID.String(), usr.Name,
	)

}

func (e EntRepository) GetUsers(ctx context.Context) (*[]user.User, error) {
	usrs, err := e.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var users []user.User
	for _, usr := range usrs {
		user, err := e.userFactory.UnmarshalUserFromDatabase(
			usr.ID.String(), usr.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}
