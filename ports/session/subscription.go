package session

import (
	"github.com/Muchogoc/semezana/domain/chat"
)

func (s *Session) addSub(channel string, sub *chat.Membership) {
	s.subscriptions.Store(channel, sub)
}

func (s *Session) getSub(channel string) *chat.Membership {
	if t, ok := s.subscriptions.Load(channel); ok {
		return t.(*chat.Membership)
	}
	return nil
}

func (s *Session) delSub(channel string) {
	s.subscriptions.Delete(channel)
}

func (s *Session) countSub() int {
	sum := 0
	s.subscriptions.Range(
		func(key, value any) bool {
			sum++
			return true
		},
	)
	return sum
}
