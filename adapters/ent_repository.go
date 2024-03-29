package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/channel"
	"github.com/Muchogoc/semezana/ent/recipient"
	"github.com/Muchogoc/semezana/ent/schema"
	"github.com/Muchogoc/semezana/ent/subscription"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

const tracerName = "github.com/Muchogoc/semezana/adapters"

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

func (e EntRepository) CreateChannel(ctx context.Context, channel *chat.Channel) error {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateChannel()")
	defer span.End()

	cid, err := uuid.Parse(channel.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil

}

func (e EntRepository) GetChannel(ctx context.Context, id string, preload bool) (*chat.Channel, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetChannel()")
	defer span.End()

	cid, err := uuid.Parse(id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	query := e.client.Channel.Query().Where(
		channel.ID(cid),
	)

	chann, err := query.Only(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	if preload {
		subs, err := chann.QuerySubscriptions().WithUser().All(ctx)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		var memberships []chat.Membership
		for _, sub := range subs {

			user, err := sub.Edges.UserOrErr()
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
				return nil, err
			}

			usr, err := e.chatFactory.UnmarshalUserFromDatabase(
				user.ID.String(), user.Name,
			)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
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
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
				return nil, err
			}

			memberships = append(memberships, *membership)

		}

		channel.SetMemberships(memberships)

	}

	return channel, nil
}

func (e EntRepository) GetChannels(ctx context.Context) (*[]chat.Channel, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetChannels()")
	defer span.End()

	chnls, err := e.client.Channel.Query().All(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
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
	_, span := otel.Tracer(tracerName).Start(ctx, "UpdateChannel()")
	defer span.End()

	return nil
}

func (e EntRepository) CreateMembership(ctx context.Context, membership *chat.Membership) error {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateMembership()")
	defer span.End()

	sid, err := uuid.Parse(membership.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	user := membership.User()
	uid, err := uuid.Parse(user.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	channel := membership.Channel()
	cid, err := uuid.Parse(channel.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func (e EntRepository) GetMembership(ctx context.Context, id string, preload bool) (*chat.Membership, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetMembership()")
	defer span.End()

	sid, err := uuid.Parse(id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	sub, err := e.client.Subscription.Query().Where(
		subscription.ID(sid),
	).WithChannel().WithUser().Only(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	channel, err := sub.Edges.ChannelOrErr()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	user, err := sub.Edges.UserOrErr()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	usr, err := e.chatFactory.UnmarshalUserFromDatabase(
		user.ID.String(), user.Name,
	)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUserChannelMembership()")
	defer span.End()

	uid, err := uuid.Parse(userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	cid, err := uuid.Parse(channelID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	sub, err := e.client.Subscription.Query().Where(
		subscription.UserID(uid),
		subscription.ChannelID(cid),
	).WithChannel().WithUser().Only(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	channel, err := sub.Edges.ChannelOrErr()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	user, err := sub.Edges.UserOrErr()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	usr, err := e.chatFactory.UnmarshalUserFromDatabase(
		user.ID.String(), user.Name,
	)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	_, span := otel.Tracer(tracerName).Start(ctx, "UpdateMembership()")
	defer span.End()

	return nil
}

func (e EntRepository) CreateMessage(ctx context.Context, message *chat.Message, userID, channelID string) error {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateMessage()")
	defer span.End()

	uid, err := uuid.Parse(userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	cid, err := uuid.Parse(channelID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	tx, err := e.client.Tx(ctx)
	if err != nil {
		err = fmt.Errorf("createMessage(): failed to start a transaction: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	author, err := tx.User.Get(ctx, uid)
	if err != nil {
		err = fmt.Errorf("createMessage(): failed to get author: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	channel, err := tx.Channel.Get(ctx, cid)
	if err != nil {
		err = fmt.Errorf("createMessage(): failed to get channel: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	sequence := channel.Sequence + 1
	_, err = tx.Channel.Update().SetSequence(sequence).Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		err = fmt.Errorf("createMessage(): failed to update channel sequence: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	mid, err := uuid.Parse(message.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
		_ = tx.Rollback()
		err = fmt.Errorf("createMessage(): failed to create message: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		err = fmt.Errorf("createMessage(): failed to commit transaction: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func (e EntRepository) CreateRecipient(ctx context.Context, resp *chat.Recipient) error {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateRecipient()")
	defer span.End()

	uid, err := uuid.Parse(resp.UserID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	mid, err := uuid.Parse(resp.MessageID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func (e EntRepository) CreateUser(ctx context.Context, user *chat.User) error {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "CreateUser()")
	defer span.End()

	uid, err := uuid.Parse(user.ID())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}

func (e EntRepository) GetUser(ctx context.Context, id string) (*chat.User, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUser()")
	defer span.End()

	uid, err := uuid.Parse(id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	usr, err := e.client.User.Get(ctx, uid)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return e.chatFactory.UnmarshalUserFromDatabase(
		usr.ID.String(), usr.Name,
	)

}

func (e EntRepository) GetUsers(ctx context.Context) (*[]chat.User, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUsers()")
	defer span.End()

	usrs, err := e.client.User.Query().All(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	var users []chat.User
	for _, usr := range usrs {
		user, err := e.chatFactory.UnmarshalUserFromDatabase(
			usr.ID.String(), usr.Name,
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}

func (e EntRepository) GetUserMemberships(ctx context.Context, userID string) (*[]chat.Membership, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetUserMemberships()")
	defer span.End()

	uid, err := uuid.Parse(userID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	subs, err := e.client.Subscription.Query().Where(
		subscription.UserID(uid),
	).WithChannel().WithUser().All(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	var memberships []chat.Membership
	for _, sub := range subs {

		channel, err := sub.Edges.ChannelOrErr()
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		user, err := sub.Edges.UserOrErr()
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		usr, err := e.chatFactory.UnmarshalUserFromDatabase(
			user.ID.String(), user.Name,
		)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
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
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
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
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}

		memberships = append(memberships, *membership)
	}

	return &memberships, nil
}

func (e EntRepository) GetAllMemberships(ctx context.Context) ([]string, error) {
	ctx, span := otel.Tracer(tracerName).Start(ctx, "GetAllMemberships()")
	defer span.End()

	var memberships []string
	subs, err := e.client.Subscription.Query().IDs(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return memberships, err
	}

	for _, sub := range subs {
		memberships = append(memberships, sub.String())
	}

	return memberships, nil
}
