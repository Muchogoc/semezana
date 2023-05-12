package chat

import (
	"time"
)

type Recipient struct {
	uid       string
	mid       string
	user      User
	status    string
	status_at time.Time
}

func (a *Recipient) User() User {
	return a.user
}

func (a *Recipient) UserID() string {
	return a.uid
}

func (a *Recipient) SetUserID(id string) {
	a.uid = id
}

func (a *Recipient) MessageID() string {
	return a.mid
}

func (a *Recipient) SetMessageID(id string) {
	a.mid = id
}

func (a *Recipient) Status() string {
	return a.status
}

func (a *Recipient) SetStatus(status string) {
	a.status = status
}

func (a *Recipient) StatusAt() time.Time {
	return a.status_at
}

func (a *Recipient) SetStatusAt(time time.Time) {
	a.status_at = time
}
