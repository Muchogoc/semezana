package models

// ObjState represents information on objects state,
// such as an indication that User or Channel is suspended/soft-deleted.
type ObjState string

const (
	// StateOK indicates normal user or channel.
	StateOK ObjState = "OK"
	// StateSuspended indicates suspended user or channel.
	StateSuspended ObjState = "SUSPENDED"
	// StateDeleted indicates soft-deleted user or channel.
	StateDeleted ObjState = "DELETED"
	// StateUndefined indicates state which has not been set explicitly.
	StateUndefined ObjState = "UNDEFINED"
)

func (o ObjState) String() string {
	return string(o)
}

// ChannelCategory is an enum of channel categories.
type ChannelCategory string

const (
	// ChannelCategoryMe is a value denoting 'me' channel.
	ChannelCategoryMe ChannelCategory = "ME"
	// ChannelCategoryFnd is a value denoting 'fnd' channel.
	ChannelCategoryFnd ChannelCategory = "FIND"
	// ChannelCategoryP2P is a a value denoting 'p2p channel.
	ChannelCategoryP2P ChannelCategory = "PEER"
	// ChannelCategoryGroup is a a value denoting group channel.
	ChannelCategoryGroup ChannelCategory = "GROUP"
	// ChannelCategoryChannel is a a value denoting group channel.
	ChannelCategoryChannel ChannelCategory = "CHANNEL"
	// ChannelCategorySystem is a constant indicating a system channel.
	ChannelCategorySystem ChannelCategory = "SYSTEM"
)

func (t ChannelCategory) String() string {
	return string(t)
}
