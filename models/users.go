package models

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound  = errors.New("entity not found")
	ErrDuplicateUser = errors.New("user with email already exists")
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	ProfilePhoto string    `json:"profilePhoto"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Verified     bool      `json:"active"`
}

type UserToken struct {
	Hash      string
	UserID    uuid.UUID
	ExpiresAt time.Time
	Scope     string
}

type UserStore interface {
	InsertUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, ID uuid.UUID) (User, error)
	GetUserByMail(ctx context.Context, email string) (User, error)
	DeleteUser(ctx context.Context, ID string) error
	InsertToken(ctx context.Context, token *UserToken) error
	GetUserForToken(ctx context.Context, tokenHash, scope, email string) (User, error)
	DeleteToken(ctx context.Context, tokenHash, scope string) error
}
