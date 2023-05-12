package chat

import (
	"time"
)

type Message struct {
	id         string
	headers    MessageHeaders
	content    MessageContent
	channel    Channel
	author     User
	recipients []Recipient
	timestamp  time.Time
}

type MessageHeaders struct{}

type MessageContent struct {
	text string
}

func (m MessageContent) Text() string {
	return m.text
}

func (m *Message) ID() string {
	return m.id
}

func (m *Message) Recipients() []Recipient {
	return m.recipients
}

func (m *Message) Headers() MessageHeaders {
	return m.headers
}

func (m *Message) Content() MessageContent {
	return m.content
}

func (m *Message) Author() User {
	return m.author
}

func (m *Message) Channel() Channel {
	return m.channel
}

func (m *Message) Timestamp() time.Time {
	return m.timestamp
}
