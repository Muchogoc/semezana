package chat

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

func (c *Channel) Messages() []Message {
	return c.messages
}
