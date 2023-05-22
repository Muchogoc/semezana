package session

import (
	"encoding/json"

	"github.com/Muchogoc/semezana/dto"
	"github.com/nsqio/go-nsq"
)

func (s *Session) HandleMessage(m *nsq.Message) error {
	var request dto.PubMessage
	if err := json.Unmarshal(m.Body, &request); err != nil {
		return err
	}

	response := s.service.ProcessPubsubMessage(s.ctx, request)
	s.queueOut(response)

	return nil
}
