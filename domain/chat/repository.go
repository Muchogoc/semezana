package chat

import (
	"context"
)

type Repository interface {
	CreateChannel(ctx context.Context, channel *Channel) error
	GetChannel(ctx context.Context, id string) (*Channel, error)
	GetChannels(ctx context.Context) (*[]Channel, error)
	UpdateChannel(
		ctx context.Context,
		cid string,
		updateFn func(h *Channel) (*Channel, error),
	) error

	CreateMembership(ctx context.Context, membership *Membership) error
	GetMembership(ctx context.Context, userID, channelID string) (*Membership, error)
	UpdateMembership(
		ctx context.Context,
		cid string,
		updateFn func(h *Membership) (*Membership, error),
	) error

	CreateMessage(ctx context.Context, message *Message, userID, channelID string) error

	CreateMessageAudience(ctx context.Context, audience *Audience) error
}
