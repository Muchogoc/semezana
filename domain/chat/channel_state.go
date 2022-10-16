package chat

// ChannelState represents information on channel state,
// such as an indication Channel is suspended/soft-deleted.
type ChannelState string

const (
	// StateOK indicates normal channel.
	StateOK ChannelState = "OK"

	// StateSuspended indicates suspended channel.
	StateSuspended ChannelState = "SUSPENDED"

	// StateDeleted indicates soft-deleted channel.
	StateDeleted ChannelState = "DELETED"

	// StateUndefined indicates state which has not been set explicitly.
	StateUndefined ChannelState = "UNDEFINED"
)

func (o ChannelState) String() string {
	return string(o)
}
