package dto

type PubSubMessageType string

type PubMessage struct {
	Sender string            `json:"sender"`
	Type   PubSubMessageType `json:"type"`
	Data   interface{}       `json:"data"`
}
