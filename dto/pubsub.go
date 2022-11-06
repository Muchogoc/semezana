package dto

type MessageType string

var (
	MessageTypeNewMembership MessageType = "NEW_MEMBERSHIP"
	MessageTypeNewMessage    MessageType = "NEW_MESSAGE"
)

type PubMessage struct {
	Sender string      `json:"sender"`
	Type   MessageType `json:"type"`
	Data   interface{} `json:"data"`
}
