package semezana

import "github.com/Muchogoc/semezana/ent"

func (s *Session) addSub(topic string, sub *ent.Subscription) {
	// s.subsLock.Lock()
	// s.subscriptions[topic] = sub
	// s.subsLock.Unlock()
	s.subscriptions.Store(topic, sub)
}

func (s *Session) getSub(topic string) *ent.Subscription {
	// s.subsLock.RLock()
	// defer s.subsLock.RUnlock()
	// return s.subscriptions[topic]
	if t, ok := s.subscriptions.Load(topic); ok {
		return t.(*ent.Subscription)
	}
	return nil
}

func (s *Session) delSub(topic string) {
	// s.subsLock.Lock()
	// delete(s.subscriptions, topic)
	// s.subsLock.Unlock()
	s.subscriptions.Delete(topic)
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
