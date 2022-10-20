package session

import (
	"encoding/json"

	"github.com/Muchogoc/semezana/dto"
	"github.com/nsqio/go-nsq"
)

func (s *Session) HandleMessage(m *nsq.Message) error {
	// s.messageLock.Lock()
	// defer s.messageLock.Unlock()

	var request dto.PubMessage
	if err := json.Unmarshal(m.Body, &request); err != nil {
		return err
	}

	s.sub <- request

	return nil
}
