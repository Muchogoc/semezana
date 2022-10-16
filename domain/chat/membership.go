package chat

import "github.com/Muchogoc/semezana/domain/user"

type Membership struct {
	id   string
	role MembershipRole

	user    user.User
	channel Channel
}

func (m *Membership) User() user.User {
	return m.user
}

func (m *Membership) Channel() Channel {
	return m.channel
}

func (m *Membership) Role() MembershipRole {
	return m.role
}

func (m *Membership) ID() string {
	return m.id
}
