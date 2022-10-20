package dto

import "time"

type ClientPayloadType string

var (
	ClientPayloadTypeHello   ClientPayloadType = "HELLO"
	ClientPayloadTypeAccount ClientPayloadType = "HELLO"
	ClientPayloadTypeLogin   ClientPayloadType = "HELLO"
	ClientPayloadTypeSub     ClientPayloadType = "HELLO"
	ClientPayloadTypeLeave   ClientPayloadType = "HELLO"
	ClientPayloadTypePublish ClientPayloadType = "HELLO"
	ClientPayloadTypeDelete  ClientPayloadType = "HELLO"
	ClientPayloadTypeNotify  ClientPayloadType = "HELLO"
)

// ClientPayload is a wrapper for client messages.
type ClientPayload struct {
	Type  ClientPayloadType `json:"type"`
	Auth  Auth              `json:"auth"`
	Extra *Extra            `json:"extra,omitempty"`

	Hello        *Hello `json:"hello,omitempty"`
	Account      *Acc   `json:"account,omitempty"`
	Login        *Login `json:"login,omitempty"`
	Subscription *Sub   `json:"subscription,omitempty"`
	Leave        *Leave `json:"leave,omitempty"`
	Publish      *Pub   `json:"publish,omitempty"`
	Delete       *Del   `json:"delete,omitempty"`
	Notify       *Note  `json:"notify,omitempty"`

	// Timestamp when this message was received by the server.
	Timestamp time.Time `json:"-"`
	// Originating session to send an acknowledgement to.
	// session *Session `json:"-"`
}

// Hello is a handshake {handshake} message.
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

// Acc is an {account} message for creating or updating a user account.
type Acc struct {
	User string `json:"user"`
}

// Login is a login {login} message.
type Login struct {
	User string `json:"user"`
}

// Sub is a subscription request {sub} message.
type Sub struct {
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// Leave is an unsubscribe {leave} request message.
type Leave struct {
}

// Pub is client's request to publish data to channel subscribers {pub}.
type Pub struct {
	Id      string                 `json:"id,omitempty"`
	User    string                 `json:"user"`
	Channel string                 `json:"channel"`
	Head    map[string]interface{} `json:"head,omitempty"`
	Content interface{}            `json:"content"`
}

// Get is a query of channel state {get}.
type Get struct{}

// Set is an update of channel state {set}.
type Set struct{}

// Del delete messages or channel {del}.
type Del struct{}

// Note is a client-generated notification for channel subscribers {note}.
type Note struct{}

// Extra is not a stand-alone message but extra data which augments the main payload.
type Extra struct{}

// Auth is used for authentication
type Auth struct {
	Token string `json:"token,omitempty"`
}
