package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

// user repository with planet scale
type UserRepository interface {
	// Authenticate user
	Authenticate(ctx context.Context, email string) (entity.User, error)
	// Get all users
	GetAllUsers() error
	// Get user by email
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	// Get user by email and password
	GetUserByEmailAndPassword(email, password string) error
	// Create new user
	CreateUser(email, password string) error
	// Update user by email
	UpdateUserByEmail(email string) error
	// Delete user by email
	DeleteUserByEmail(email string) error
}
