package chat

import (
	"github.com/Muchogoc/semezana/domain/user"
)

type Membership struct {
	id      string
	role    MembershipRole
	status  MembershipStatus
	user    user.User
	channel Channel
}

func (m *Membership) User() user.User {
	return m.user
}

func (m *Membership) SetUser(user user.User) {
	m.user = user
}

func (m *Membership) Channel() Channel {
	return m.channel
}

func (m *Membership) SetChannel(channel Channel) {
	m.channel = channel
}

func (m *Membership) Role() MembershipRole {
	return m.role
}

func (m *Membership) SetRole(role MembershipRole) {
	m.role = role
}

func (m *Membership) ID() string {
	return m.id
}

func (m *Membership) SetID(id string) {
	m.id = id
}

func (m *Membership) Status() MembershipStatus {
	return m.status
}

func (m *Membership) SetStatus(status MembershipStatus) {
	m.status = status
}

func (m *Membership) CanReceiveMessage() bool {
	return m.Status() == MembershipStatusOk
}
