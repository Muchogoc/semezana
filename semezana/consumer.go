package semezana

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/semezana/dto"
	"github.com/nsqio/go-nsq"
)

type PubMessage struct {
	Sender  string       `json:"sender,omitempty"`
	Type    string       `json:"type,omitempty"`
	Message *ent.Message `json:"message,omitempty"`
}

func (m PubMessage) Marshal() []byte {
	marshalled, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	return marshalled
}

func (s *Session) nsqConsumer(ctx context.Context, subscription *ent.Subscription) {
	config := nsq.NewConfig()

	nsqTopic := subscription.ID
	// ephemeral channels disappear after last client disconnects.
	nsqChannel := fmt.Sprintf("%s#ephemeral", s.sid)

	//Creating the consumer
	consumer, err := nsq.NewConsumer(nsqTopic.String(), nsqChannel, config)
	if err != nil {
		log.Fatal(err)
	}

	if DEBUG {
		consumer.SetLogger(log.Default(), nsq.LogLevelDebug)
	}

	consumer.AddHandler(s)

	//Creating the Producer using NSQ lookup Address
	consumer.ConnectToNSQLookupd(NSQ_LOOKUP_ADDRESS)
	<-s.stop

	consumer.Stop()
}

func (s *Session) HandleMessage(m *nsq.Message) error {
	// s.messageLock.Lock()
	// defer s.messageLock.Unlock()

	var request PubMessage
	if err := json.Unmarshal(m.Body, &request); err != nil {
		return err
	}

	if request.Type == "message.new" {
		response := &dto.ServerComMessage{
			Data: &dto.MsgServerData{
				Topic:     request.Message.ID.String(),
				From:      request.Sender,
				Timestamp: request.Message.CreatedAt,
				SeqId:     request.Message.Sequence,
				Head:      request.Message.Header,
				Content:   request.Message.Content["content"],
			},
			Control: &dto.MsgServerCtrl{
				Code: http.StatusOK,
			},
		}

		s.queueOut(response)
	}

	return nil
}
