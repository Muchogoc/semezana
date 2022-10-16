package chat

type MembershipRole string

const (
	MembershipRoleAdmin     MembershipRole = "ADMIN"
	MembershipRoleModerator MembershipRole = "MODERATOR"
	MembershipRoleOwner     MembershipRole = "OWNER"
)

func (m MembershipRole) String() string {
	return string(m)
}
