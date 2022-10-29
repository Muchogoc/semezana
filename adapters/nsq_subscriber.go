package adapters

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Muchogoc/semezana/internal/utils"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

var (
	DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
)

type NSQSubscriber struct {
	lookupAddress string
}

func NewNSQSubscriber(address string) *NSQSubscriber {
	return &NSQSubscriber{
		lookupAddress: address,
	}
}

func (n NSQSubscriber) CreateSessionSubscriber(ctx context.Context, membershipID string) {
	session, err := utils.SessionFromContext(ctx)
	if err != nil {
		logrus.Error(err)
		return
	}

	config := nsq.NewConfig()

	// ephemeral channels disappear after last client disconnects.
	nsqChannel := fmt.Sprintf("%s#ephemeral", session.ID())

	//Creating the consumer
	consumer, err := nsq.NewConsumer(membershipID, nsqChannel, config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Stop()

	if DEBUG {
		consumer.SetLogger(log.Default(), nsq.LogLevelDebug)
	}

	consumer.AddHandler(session)

	//Creating the Producer using NSQ lookup Address
	consumer.ConnectToNSQLookupd(n.lookupAddress)

	stop := session.StopChan()
	<-stop

}
