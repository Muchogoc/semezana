package chat

import (
	"fmt"

	"github.com/Muchogoc/semezana/domain/user"
	"github.com/google/uuid"
)

type Channel struct {
	id          string
	description string
	name        string
	state       ChannelState
	category    ChannelCategory

	messages    []Message
	memberships []Membership
}

func (c *Channel) ID() string {
	return c.id
}

func (c *Channel) SetID(id string) {
	c.id = id
}

func (c *Channel) Name() string {
	return c.name
}

func (c *Channel) SetName(name string) {
	c.name = name
}

func (c *Channel) Description() string {
	return c.description
}

func (c *Channel) SetDescription(description string) {
	c.description = description
}

func (c *Channel) Category() ChannelCategory {
	return c.category
}

func (c *Channel) SetCategory(category ChannelCategory) {
	c.category = category
}

func (c *Channel) State() ChannelState {
	return c.state
}

func (c *Channel) SetState(state ChannelState) {
	c.state = state
}

func (c *Channel) Memberships() []Membership {
	return c.memberships
}

func (c *Channel) SetMemberships(memberships []Membership) {
	c.memberships = memberships
}

func (c *Channel) Messages() []Message {
	return c.messages
}

func (c *Channel) SetMessages(messages []Message) {
	c.messages = messages
}

func (c *Channel) ValidateMembership(uid string) (user.User, error) {
	var exists bool
	var current user.User

	for _, membership := range c.Memberships() {
		user := membership.User()
		if user.ID() == uid {
			current = user
			exists = true
			break
		}
	}

	if !exists {
		return current, fmt.Errorf("user is not a member of channel:%s", c.name)
	}

	return current, nil
}

// New Message receives a text. It returns the new message and the memberships that should receive it
func (c *Channel) NewMessage(text string, senderID string) (*Message, *[]Membership, error) {
	user, err := c.ValidateMembership(senderID)
	if err != nil {
		return nil, nil, err
	}

	var recipients []Membership

	for _, membership := range c.Memberships() {
		if membership.CanReceiveMessage() {
			recipients = append(recipients, membership)
		}
	}

	message := Message{
		id:      uuid.NewString(),
		headers: MessageHeaders{},
		content: MessageContent{
			text: text,
		},
		channel: *c,
		author:  user,
	}

	return &message, &recipients, nil
}
