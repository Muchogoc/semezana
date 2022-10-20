package adapters

import (
	"context"
	"time"

	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/schema"
	"github.com/Muchogoc/semezana/ent/subscription"
	"github.com/google/uuid"
)

type EntRepository struct {
	client      *ent.Client
	chatFactory chat.Factory
}

func NewEntRepository(client *ent.Client, chatFactory chat.Factory) *EntRepository {
	return &EntRepository{
		client:      client,
		chatFactory: chatFactory,
	}
}

func (e EntRepository) CreateChannel(ctx context.Context, channel chat.Channel) error {
	_, err := e.client.Channel.Create().
		SetName(channel.Name()).
		SetDescription(channel.Description()).
		SetType(channel.Category().String()).
		SetState(channel.State().String()).
		SetStateAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}
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

func (e EntRepository) UpdateChannel(
	ctx context.Context,
	id string,
	updateFn func(h *chat.Channel) (*chat.Channel, error),
) error {

	return nil
}

func (e EntRepository) CreateMembership(ctx context.Context, membership chat.Membership) error {
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

func (e EntRepository) CreateMessage(ctx context.Context, message chat.Message, userID, channelID string) error {
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

func (e EntRepository) CreateMessageAudience(ctx context.Context, audience chat.Audience) error {
	return nil
}
