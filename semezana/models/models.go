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

type Topic struct {
	Abstract

	// State of the topic: normal (ok), suspended, deleted
	State   ObjState
	StateAt *time.Time
}

// A subscription to a topic
type Subscription struct {
	Abstract

	// User who has relationship with the topic
	User string
	// Topic subscribed to
	Topic string
}

// Represents a single message sent between users
type Message struct {
	Abstract
	SeqID     int
	Topic     string
	From      string
	Content   interface{}
	DeletedAt *time.Time
}
