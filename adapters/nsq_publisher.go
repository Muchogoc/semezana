package adapters

import (
	"context"
	"encoding/json"

	"github.com/Muchogoc/semezana/dto"
	"github.com/nsqio/go-nsq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
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
	_, span := otel.Tracer(tracerName).Start(ctx, "PublishToMembership()")
	defer span.End()

	marshalled, err := json.Marshal(input)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	err = p.producer.Publish(membershipID, marshalled)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return err
	}

	return nil
}
