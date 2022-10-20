package chat

type ChannelConfig struct {
	typingEvents     bool
	reactions        bool
	replies          bool
	threads          bool
	uploads          bool
	flag             bool
	maxMessageLength int
}

func (c *ChannelConfig) TypingEvents() bool {
	return c.typingEvents
}

func (c *ChannelConfig) Reactions() bool {
	return c.reactions
}

func (c *ChannelConfig) Replies() bool {
	return c.replies
}

func (c *ChannelConfig) Threads() bool {
	return c.threads
}

func (c *ChannelConfig) Uploads() bool {
	return c.uploads
}

func (c *ChannelConfig) Flag() bool {
	return c.flag
}

func (c *ChannelConfig) MaxMessageLength() int {
	return c.maxMessageLength
}
