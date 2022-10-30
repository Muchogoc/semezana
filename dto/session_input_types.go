package dto

import "time"

type ClientPayloadType string

var (
	ClientPayloadTypeHello   ClientPayloadType = "HELLO"
	ClientPayloadTypePublish ClientPayloadType = "PUBLISH"
	ClientPayloadTypeNotify  ClientPayloadType = "NOTIFY"
)

// ClientPayload is a wrapper for client messages.
type ClientPayload struct {
	Type  ClientPayloadType `json:"type"`
	Auth  Auth              `json:"auth"`
	Extra *Extra            `json:"extra,omitempty"`

	Hello   *Hello `json:"hello,omitempty"`
	Publish *Pub   `json:"publish,omitempty"`
	Notify  *Note  `json:"notify,omitempty"`

	// Timestamp when this message was received by the server.
	Timestamp time.Time `json:"-"`
}

// Hello is a handshake message sent immediately after the connection is set up.
type Hello struct {
	// User agent
	UserAgent string `json:"userAgent,omitempty"`
	// Protocol version, i.e. "0.13"
	Version string `json:"version,omitempty"`
	// Client's unique device ID
	DeviceID string `json:"deviceID,omitempty"`
	// ISO 639-1 human language of the connected device
	Language string `json:"language,omitempty"`
	// Platform code: ios, android, web.
	Platform string `json:"platform,omitempty"`
}

// Pub is client's request to publish data to channel subscribers {pub}.
type Pub struct {
	Id      string                 `json:"id,omitempty"`
	User    string                 `json:"user"`
	Channel string                 `json:"channel"`
	Head    map[string]interface{} `json:"head,omitempty"`
	Content interface{}            `json:"content"`
}

// Note is a client-generated notification for channel subscribers {note}.
type Note struct{}

// Extra is not a stand-alone message but extra data which augments the main payload.
type Extra struct{}

// Auth is used for authentication
type Auth struct {
	Token string `json:"token,omitempty"`
}
