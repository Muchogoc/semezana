package chat

import "github.com/Muchogoc/semezana/domain/user"

type Message struct {
	id       string
	headers  MessageHeaders
	content  MessageContent
	channel  Channel
	author   user.User
	audience []Audience
}

type MessageHeaders struct{}

type MessageContent struct {
	text string
}

func (m *MessageContent) Text() string {
	return m.text
}

func (m *Message) ID() string {
	return m.id
}

func (m *Message) Audience() []Audience {
	return m.audience
}

func (m *Message) Headers() MessageHeaders {
	return m.headers
}

func (m *Message) Content() MessageContent {
	return m.content
}

func (m *Message) Author() user.User {
	return m.author
}

func (m *Message) Channel() Channel {
	return m.channel
}
