package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/domain/user"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/channel"
	"github.com/Muchogoc/semezana/ent/recipient"
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

	_, err = e.client.Channel.Create().
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

	return nil

}

func (e EntRepository) GetChannel(ctx context.Context, id string, preload bool) (*chat.Channel, error) {
	cid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	query := e.client.Channel.Query().Where(
		channel.ID(cid),
	)

	chann, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}

	channel, err := e.chatFactory.UnmarshalChannelFromDatabase(
		chann.ID.String(),
		chann.Description,
		chann.Name,
		chat.ChannelState(chann.State),
		chat.ChannelCategory(chann.Type),
	)
	if err != nil {
		return nil, err
	}

	if preload {
		subs, err := chann.QuerySubscriptions().WithUser().All(ctx)
		if err != nil {
			return nil, err
		}

		var memberships []chat.Membership
		for _, sub := range subs {

			user, err := sub.Edges.UserOrErr()
			if err != nil {
				return nil, err
			}

			usr, err := e.userFactory.UnmarshalUserFromDatabase(
				user.ID.String(), user.Name,
			)
			if err != nil {
				return nil, err
			}

			membership, err := e.chatFactory.UnmarshalMembershipFromDatabase(
				sub.ID.String(),
				chat.MembershipRole(sub.Role),
				chat.MembershipStatus(sub.Status),
				*channel,
				*usr,
			)
			if err != nil {
				return nil, err
			}

			memberships = append(memberships, *membership)

		}

		channel.SetMemberships(memberships)

	}

	return channel, nil
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
	sid, err := uuid.Parse(membership.ID())
	if err != nil {
		return err
	}

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
		SetID(sid).
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

func (e EntRepository) GetMembership(ctx context.Context, id string, preload bool) (*chat.Membership, error) {
	sid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	sub, err := e.client.Subscription.Query().Where(
		subscription.ID(sid),
	).WithChannel().WithUser().Only(ctx)
	if err != nil {
		return nil, err
	}

	channel, err := sub.Edges.ChannelOrErr()
	if err != nil {
		return nil, err
	}

	user, err := sub.Edges.UserOrErr()
	if err != nil {
		return nil, err
	}

	usr, err := e.userFactory.UnmarshalUserFromDatabase(
		user.ID.String(), user.Name,
	)
	if err != nil {
		return nil, err
	}

	chn, err := e.chatFactory.UnmarshalChannelFromDatabase(
		channel.ID.String(),
		channel.Description,
		channel.Name,
		chat.ChannelState(channel.State),
		chat.ChannelCategory(channel.Type),
	)
	if err != nil {
		return nil, err
	}

	return e.chatFactory.UnmarshalMembershipFromDatabase(
		sub.ID.String(),
		chat.MembershipRole(sub.Role),
		chat.MembershipStatus(sub.Status),
		*chn,
		*usr,
	)

}

func (e EntRepository) GetUserChannelMembership(ctx context.Context, userID, channelID string) (*chat.Membership, error) {
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
	).WithChannel().WithUser().Only(ctx)
	if err != nil {
		return nil, err
	}

	channel, err := sub.Edges.ChannelOrErr()
	if err != nil {
		return nil, err
	}

	user, err := sub.Edges.UserOrErr()
	if err != nil {
		return nil, err
	}

	usr, err := e.userFactory.UnmarshalUserFromDatabase(
		user.ID.String(), user.Name,
	)
	if err != nil {
		return nil, err
	}

	chn, err := e.chatFactory.UnmarshalChannelFromDatabase(
		channel.ID.String(),
		channel.Description,
		channel.Name,
		chat.ChannelState(channel.State),
		chat.ChannelCategory(channel.Type),
	)
	if err != nil {
		return nil, err
	}

	return e.chatFactory.UnmarshalMembershipFromDatabase(
		sub.ID.String(),
		chat.MembershipRole(sub.Role),
		chat.MembershipStatus(sub.Status),
		*chn,
		*usr,
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
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	cid, err := uuid.Parse(channelID)
	if err != nil {
		return err
	}

	tx, err := e.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("createMessage(): failed to start a transaction: %w", err)
	}

	author, err := tx.User.Get(ctx, uid)
	if err != nil {
		return fmt.Errorf("createMessage(): failed to get author: %w", err)
	}

	channel, err := tx.Channel.Get(ctx, cid)
	if err != nil {
		return fmt.Errorf("createMessage(): failed to get channel: %w", err)
	}

	sequence := channel.Sequence + 1
	_, err = tx.Channel.Update().SetSequence(sequence).Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("createMessage(): failed to update channel sequence: %w", err)
	}

	mid, err := uuid.Parse(message.ID())
	if err != nil {
		return err
	}

	content := message.Content()
	_, err = tx.Message.Create().
		SetID(mid).
		SetAuthor(author).
		SetChannel(channel).
		SetHeader(schema.MessageHeaders{}).
		SetContent(schema.MessageContent{
			Text: content.Text(),
		}).
		SetSequence(sequence).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("createMessage(): failed to create message: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("createMessage(): failed to commit transaction: %w", err)
	}

	return nil
}

func (e EntRepository) CreateRecipient(ctx context.Context, resp *chat.Recipient) error {
	uid, err := uuid.Parse(resp.UserID())
	if err != nil {
		return err
	}

	mid, err := uuid.Parse(resp.MessageID())
	if err != nil {
		return err
	}

	_, err = e.client.Recipient.
		Create().
		SetMessageID(mid).
		SetUserID(uid).
		SetStatus(recipient.Status(resp.Status())).
		SetStatusAt(resp.StatusAt()).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (e EntRepository) CreateUser(ctx context.Context, user *user.User) error {
	uid, err := uuid.Parse(user.ID())
	if err != nil {
		return err
	}

	_, err = e.client.User.
		Create().
		SetID(uid).
		SetName(user.Name()).
		SetState("OK").
		SetStateAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

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

func (e EntRepository) GetUserMemberships(ctx context.Context, userID string) (*[]chat.Membership, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	subs, err := e.client.Subscription.Query().Where(
		subscription.UserID(uid),
	).WithChannel().WithUser().All(ctx)
	if err != nil {
		return nil, err
	}

	var memberships []chat.Membership
	for _, sub := range subs {

		channel, err := sub.Edges.ChannelOrErr()
		if err != nil {
			return nil, err
		}

		user, err := sub.Edges.UserOrErr()
		if err != nil {
			return nil, err
		}

		usr, err := e.userFactory.UnmarshalUserFromDatabase(
			user.ID.String(), user.Name,
		)
		if err != nil {
			return nil, err
		}

		chn, err := e.chatFactory.UnmarshalChannelFromDatabase(
			channel.ID.String(),
			channel.Description,
			channel.Name,
			chat.ChannelState(channel.State),
			chat.ChannelCategory(channel.Type),
		)
		if err != nil {
			return nil, err
		}

		membership, err := e.chatFactory.UnmarshalMembershipFromDatabase(
			sub.ID.String(),
			chat.MembershipRole(sub.Role),
			chat.MembershipStatus(sub.Status),
			*chn,
			*usr,
		)
		if err != nil {
			return nil, err
		}

		memberships = append(memberships, *membership)
	}

	return &memberships, nil
}

func (e EntRepository) GetAllMemberships(ctx context.Context) ([]string, error) {
	var memberships []string
	subs, err := e.client.Subscription.Query().IDs(ctx)
	if err != nil {
		return memberships, err
	}

	for _, sub := range subs {
		memberships = append(memberships, sub.String())
	}

	return memberships, nil
}
