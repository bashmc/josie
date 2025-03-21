package models

import (
	"context"
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Active       bool      `json:"active"`
}

type UserStore interface {
	InsertUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, ID string) (User, error)
	DeleteUser(ctx context.Context, ID string) error
}
