package dto

import (
	"time"

	"github.com/Muchogoc/semezana/ent/schema"
)

type ServerResponseType string

var (
	ServerResponseTypeControl  ServerResponseType = "CONTROL"
	ServerResponseTypeData     ServerResponseType = "DATA"
	ServerResponseTypeMeta     ServerResponseType = "META"
	ServerResponseTypePresence ServerResponseType = "PRESENCE"
	ServerResponseTypeInfo     ServerResponseType = "INFO"
)

// ServerResponse is a wrapper for server-side messages.
type ServerResponse struct {
	Type ServerResponseType `json:"type"`

	Control  *Ctrl `json:"control,omitempty"`
	Data     *Data `json:"data,omitempty"`
	Meta     *Meta `json:"meta,omitempty"`
	Presence *Pres `json:"presence,omitempty"`
	Info     *Info `json:"info,omitempty"`
}

// Ctrl is a server control message {ctrl}.
type Ctrl struct {
	Code       int         `json:"code,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
	Message    string      `json:"message,omitempty"`
	Line       int
}

// Data is a server {data} message.
type Data struct {
	Head      schema.MessageHeaders `json:"head,omitempty"`
	Channel   string                `json:"channel"`
	From      string                `json:"from,omitempty"`
	Timestamp time.Time             `json:"timestamp"`
	Sequence  int                   `json:"sequence"`
	Content   schema.MessageContent `json:"content"`
}

// Pres is presence notification {pres} (authoritative update).
type Pres struct{}

// Meta is a channel metadata {meta} update.
type Meta struct{}

// Info is the server-side copy of MsgClientNote with From and optionally Src added (non-authoritative).
type Info struct{}
