package main

type Subscription struct{}

func (s *Session) addSub(topic string, sub *Subscription) {
	s.subsLock.Lock()
	s.subscriptions[topic] = sub
	s.subsLock.Unlock()
}

func (s *Session) getSub(topic string) *Subscription {
	s.subsLock.RLock()
	defer s.subsLock.RUnlock()

	return s.subscriptions[topic]
}

func (s *Session) delSub(topic string) {
	s.subsLock.Lock()
	delete(s.subscriptions, topic)
	s.subsLock.Unlock()
}

func (s *Session) countSub() int {
	return len(s.subscriptions)
}
