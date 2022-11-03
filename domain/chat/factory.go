package chat

import "github.com/Muchogoc/semezana/domain/user"

type Factory struct{}

func NewFactory() (Factory, error) {
	return Factory{}, nil
}

func (f Factory) NewChannel() Channel {
	return Channel{}
}

func (f Factory) NewMessage() Message {
	return Message{}
}

func (f Factory) UnmarshalChannelFromDatabase(
	id string,
	description string,
	name string,
	state ChannelState,
	category ChannelCategory,
) (*Channel, error) {
	return &Channel{
		id:          id,
		description: description,
		name:        name,
		state:       state,
		category:    category,
	}, nil
}

func (f Factory) UnmarshalMembershipFromDatabase(
	id string,
	role MembershipRole,
	channel Channel,
	user user.User,
) (*Membership, error) {
	return &Membership{
		id:      id,
		role:    role,
		channel: channel,
		user:    user,
	}, nil
}
