package user

type User struct {
	id string
}

func (u *User) ID() string {
	return u.id
}
