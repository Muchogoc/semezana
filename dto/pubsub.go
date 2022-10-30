package dto

type PubSubMessageType string

var (
	PubSubMessageTypeCreate     PubSubMessageType = "CREATE_TOPIC"
	PubSubMessageTypeNewMessage PubSubMessageType = "RECEIVED_MESSAGE"
)

type PubMessage struct {
	Sender string            `json:"sender"`
	Type   PubSubMessageType `json:"type"`
	Data   interface{}       `json:"data"`
}
