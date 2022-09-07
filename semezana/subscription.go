package semezana

import "github.com/Muchogoc/semezana/ent"

func (s *Session) addSub(channel string, sub *ent.Subscription) {
	// s.subsLock.Lock()
	// s.subscriptions[channel] = sub
	// s.subsLock.Unlock()
	s.subscriptions.Store(channel, sub)
}

func (s *Session) getSub(channel string) *ent.Subscription {
	// s.subsLock.RLock()
	// defer s.subsLock.RUnlock()
	// return s.subscriptions[channel]
	if t, ok := s.subscriptions.Load(channel); ok {
		return t.(*ent.Subscription)
	}
	return nil
}

func (s *Session) delSub(channel string) {
	// s.subsLock.Lock()
	// delete(s.subscriptions, channel)
	// s.subsLock.Unlock()
	s.subscriptions.Delete(channel)
}

func (s *Session) countSub() int {
	// return len(s.subscriptions)
	sum := 0
	s.subscriptions.Range(
		func(key, value any) bool {
			sum++
			return true
		},
	)
	return sum
}
