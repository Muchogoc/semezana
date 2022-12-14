package utils

import (
	"context"
	"fmt"

	"github.com/Muchogoc/semezana/dto"
	"github.com/Muchogoc/semezana/ports/session"
)

func SessionFromContext(ctx context.Context) (*session.Session, error) {
	value := ctx.Value(dto.ContextKeySession)
	if value == nil {
		return nil, fmt.Errorf("no session in context")
	}

	session, ok := value.(*session.Session)
	if !ok {
		return nil, fmt.Errorf("invalid session type in context")
	}

	return session, nil
}
