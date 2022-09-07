package models

import "time"

type Abstract struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User is a representation of a user record
type User struct {
	Abstract
}

type Channel struct {
	Abstract

	// State of the channel: normal (ok), suspended, deleted
	State   ObjState
	StateAt *time.Time
}

// A subscription to a channel
type Subscription struct {
	Abstract

	// User who has relationship with the channel
	User string
	// Channel subscribed to
	Channel string
}

// Represents a single message sent between users
type MessageContent struct {
	Text string `json:"text"`
}
