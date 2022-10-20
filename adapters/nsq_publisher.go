package adapters

import (
	"context"
	"encoding/json"

	"github.com/Muchogoc/semezana/dto"
	"github.com/nsqio/go-nsq"
)

type NSQPublisher struct {
	producer *nsq.Producer
}

func NewNSQPublisher(producer *nsq.Producer) *NSQPublisher {
	return &NSQPublisher{
		producer: producer,
	}
}

func (p NSQPublisher) PublishToMembership(ctx context.Context, membershipID string, input dto.PubMessage) error {
	marshalled, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = p.producer.Publish(membershipID, marshalled)
	if err != nil {
		return err
	}

	return nil
}
