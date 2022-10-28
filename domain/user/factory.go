package user

type Factory struct{}

func NewFactory() (Factory, error) {
	return Factory{}, nil
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
