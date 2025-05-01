package services

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/shcmd/split/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store models.UserStore
}

func NewUserService(us models.UserStore) *UserService {
	return &UserService{
		store: us,
	}
}

// CreateUser creates a new user with the given details
func (s *UserService) CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		return nil, err
	}
	now := time.Now().UTC()
	user := &models.User{
		ID:           uuid.New().String(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    now,
		UpdatedAt:    now,
		Active:       true,
	}

	if err := s.store.InsertUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user's details
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now().UTC()
	return s.store.UpdateUser(ctx, user)
}

// FetchUser retrieves a user by ID or email
func (s *UserService) FetchUser(ctx context.Context, identifier string) (*models.User, error) {
	user, err := s.store.GetUser(ctx, identifier)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser removes a user from the system
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.store.DeleteUser(ctx, id)
}
