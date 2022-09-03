package dto

import "time"

// MsgServerCtrl is a server control message {ctrl}.
type MsgServerCtrl struct {
	Code       int         `json:"code,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
	Message    string      `json:"message,omitempty"`
}

// MsgServerData is a server {data} message.
type MsgServerData struct {
	Topic     string                 `json:"topic"`
	From      string                 `json:"from,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
	SeqId     int                    `json:"sequence"`
	Head      map[string]interface{} `json:"head,omitempty"`
	Content   interface{}            `json:"content"`
}

// MsgServerPres is presence notification {pres} (authoritative update).
type MsgServerPres struct{}

// MsgServerMeta is a topic metadata {meta} update.
type MsgServerMeta struct{}

// MsgServerInfo is the server-side copy of MsgClientNote with From and optionally Src added (non-authoritative).
type MsgServerInfo struct{}

// ServerComMessage is a wrapper for server-side messages.
type ServerComMessage struct {
	Control  *MsgServerCtrl `json:"control,omitempty"`
	Data     *MsgServerData `json:"data,omitempty"`
	Meta     *MsgServerMeta `json:"meta,omitempty"`
	Presence *MsgServerPres `json:"presence,omitempty"`
	Info     *MsgServerInfo `json:"info,omitempty"`
}

// type WSOutput struct {
// 	Data  *ServerComMessage `json:"data,omitempty"`
// 	Error string            `json:"error,omitempty"`
// }
