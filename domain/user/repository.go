package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context) (*[]User, error)
}
