package app

import (
	"github.com/Muchogoc/semezana/domain/chat"
	"github.com/Muchogoc/semezana/domain/user"
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

	return ChatService{
		chatRepo:   chatRepo,
		userRepo:   userRepo,
		publisher:  publisher,
		subscriber: subscriber,
	}
}
