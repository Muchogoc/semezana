package app

import (
	"context"

	"github.com/Muchogoc/semezana/dto"
)

type Publish interface {
	PublishToMembership(ctx context.Context, membershipID string, input dto.PubMessage) error
}

type Subscribe interface {
	CreateSessionSubscriber(ctx context.Context, membershipID string)
}
