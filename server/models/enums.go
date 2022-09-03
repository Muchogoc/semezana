package models

// ObjState represents information on objects state,
// such as an indication that User or Topic is suspended/soft-deleted.
type ObjState string

const (
	// StateOK indicates normal user or topic.
	StateOK ObjState = "OK"
	// StateSuspended indicates suspended user or topic.
	StateSuspended ObjState = "SUSPENDED"
	// StateDeleted indicates soft-deleted user or topic.
	StateDeleted ObjState = "DELETED"
	// StateUndefined indicates state which has not been set explicitly.
	StateUndefined ObjState = "UNDEFINED"
)

func (o ObjState) String() string {
	return string(o)
}

// TopicCategory is an enum of topic categories.
type TopicCategory string

const (
	// TopicCategoryMe is a value denoting 'me' topic.
	TopicCategoryMe TopicCategory = "ME"
	// TopicCategoryFnd is a value denoting 'fnd' topic.
	TopicCategoryFnd TopicCategory = "FIND"
	// TopicCategoryP2P is a a value denoting 'p2p topic.
	TopicCategoryP2P TopicCategory = "PEER"
	// TopicCategoryGroup is a a value denoting group topic.
	TopicCategoryGroup TopicCategory = "GROUP"
	// TopicCategoryChannel is a a value denoting group topic.
	TopicCategoryChannel TopicCategory = "CHANNEL"
	// TopicCategorySystem is a constant indicating a system topic.
	TopicCategorySystem TopicCategory = "SYSTEM"
)
