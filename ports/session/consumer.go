package session

import (
	"context"
	"encoding/json"

	"github.com/Muchogoc/semezana/dto"
	"github.com/nsqio/go-nsq"
)

func (s *Session) HandleMessage(m *nsq.Message) error {
	var request dto.PubMessage
	if err := json.Unmarshal(m.Body, &request); err != nil {
		return err
	}

	s.sub <- request

	return nil
}

func (s *Session) SubscriptionListener() {
	for {
		select {
		case msg := <-s.sub:
			response := s.service.ProcessPubsubMessage(context.Background(), msg)
			s.queueOut(response)

		case <-s.stop:
			return
		}
	}
}
