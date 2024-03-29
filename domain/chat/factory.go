package chat

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
	status MembershipStatus,
	channel Channel,
	user User,
) (*Membership, error) {
	return &Membership{
		id:      id,
		role:    role,
		status:  status,
		channel: channel,
		user:    user,
	}, nil
}

func (f Factory) NewUser(name string) User {
	return User{
		name: name,
	}
}

func (f Factory) UnmarshalUserFromDatabase(
	id string,
	name string,
) (*User, error) {
	return &User{
		id:   id,
		name: name,
	}, nil
}
