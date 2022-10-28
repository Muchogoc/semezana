package user

type User struct {
	id   string
	name string
}

func (u *User) ID() string {
	return u.id
}

func (u *User) SetID(id string) {
	u.id = id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
