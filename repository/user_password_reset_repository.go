package repository

import "context"

type UserPasswordResetRepository interface {
	// Create User Password Reset
	CreateUserPasswordReset(ctx context.Context, email string) (string, error)
	// Get User Password Reset
	GetUserPasswordResetToEmail(ctx context.Context, token string) (string, error)
	// Delete User Password Reset
	DeleteUserPasswordReset(ctx context.Context, token string) error
}
