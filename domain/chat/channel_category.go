package chat

// ChannelCategory is an enum of channel categories.
type ChannelCategory string

const (
	// ChannelCategoryP2P is a a value denoting 'p2p channel.
	// PEER is only between two people
	ChannelCategoryP2P ChannelCategory = "PEER"

	// ChannelCategoryGroup is a a value denoting group channel.
	// Group is where users can send messages to each other
	ChannelCategoryGroup ChannelCategory = "GROUP"

	// ChannelCategoryChannel is a a value denoting group channel.
	// Channel is where only approved users can send messages
	ChannelCategoryChannel ChannelCategory = "CHANNEL"
)

func (t ChannelCategory) String() string {
	return string(t)
}
