package dto

import "time"

// MsgClientHi is a handshake {handshake} message.
type MsgClientHi struct {
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

// MsgClientAcc is an {account} message for creating or updating a user account.
type MsgClientAcc struct {
	User string `json:"user"`
}

// MsgClientLogin is a login {login} message.
type MsgClientLogin struct {
	User string `json:"user"`
}

// MsgClientSub is a subscription request {sub} message.
type MsgClientSub struct {
	User  string `json:"user"`
	Topic string `json:"topic"`
}

// MsgClientLeave is an unsubscribe {leave} request message.
type MsgClientLeave struct {
}

// MsgClientPub is client's request to publish data to topic subscribers {pub}.
type MsgClientPub struct {
	Id      string                 `json:"id,omitempty"`
	User    string                 `json:"user"`
	Topic   string                 `json:"topic"`
	Head    map[string]interface{} `json:"head,omitempty"`
	Content interface{}            `json:"content"`
}

// MsgClientGet is a query of topic state {get}.
type MsgClientGet struct{}

// MsgClientSet is an update of topic state {set}.
type MsgClientSet struct{}

// MsgClientDel delete messages or topic {del}.
type MsgClientDel struct{}

// MsgClientNote is a client-generated notification for topic subscribers {note}.
type MsgClientNote struct{}

// MsgClientExtra is not a stand-alone message but extra data which augments the main payload.
type MsgClientExtra struct{}

// Auth is used for authentication
type Auth struct {
	Token string `json:"token,omitempty"`
}

// ClientComMessage is a wrapper for client messages.
type ClientComMessage struct {
	Type         string          `json:"type"`
	Auth         string          `json:"auth,omitempty"`
	Hi           *MsgClientHi    `json:"hi,omitempty"`
	Account      *MsgClientAcc   `json:"account,omitempty"`
	Login        *MsgClientLogin `json:"login,omitempty"`
	Subscription *MsgClientSub   `json:"subscription,omitempty"`
	Leave        *MsgClientLeave `json:"leave,omitempty"`
	Publish      *MsgClientPub   `json:"publish,omitempty"`
	Delete       *MsgClientDel   `json:"delete,omitempty"`
	Notify       *MsgClientNote  `json:"notify,omitempty"`
	Extra        *MsgClientExtra `json:"extra,omitempty"`

	// // Message ID
	// ID string `json:"-"`
	// // Sender's UserId as string.
	// SenderID string `json:"-"`

	// //Topics
	// Originator string `json:"-"` // Original Topic
	Receiver string `json:"-"` // Destination Topic

	// Timestamp when this message was received by the server.
	Timestamp time.Time `json:"-"`
	// Originating session to send an acknowledgement to.
	// session *Session `json:"-"`
}
