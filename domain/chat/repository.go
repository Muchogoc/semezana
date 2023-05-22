package chat

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context) (*[]User, error)

	CreateChannel(ctx context.Context, channel *Channel) error
	GetChannel(ctx context.Context, id string, preload bool) (*Channel, error)
	GetChannels(ctx context.Context) (*[]Channel, error)
	UpdateChannel(
		ctx context.Context,
		cid string,
		updateFn func(h *Channel) (*Channel, error),
	) error

	GetAllMemberships(ctx context.Context) ([]string, error)
	CreateMembership(ctx context.Context, membership *Membership) error
	GetMembership(ctx context.Context, id string, preload bool) (*Membership, error)
	GetUserMemberships(ctx context.Context, userID string) (*[]Membership, error)
	GetUserChannelMembership(ctx context.Context, userID, channelID string) (*Membership, error)
	UpdateMembership(
		ctx context.Context,
		cid string,
		updateFn func(h *Membership) (*Membership, error),
	) error

	CreateMessage(ctx context.Context, message *Message, userID, channelID string) error
	CreateRecipient(ctx context.Context, resp *Recipient) error
}
