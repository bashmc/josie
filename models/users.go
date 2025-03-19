package models

import "context"

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createAt"`
}

type UserStore interface {
	CreatUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, Id string) (User, error)
	DeleteUser(ctx context.Context, Id string) error
}
