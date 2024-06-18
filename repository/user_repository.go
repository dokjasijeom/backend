package repository

import (
	"context"
	"github.com/dokjasijeom/backend/common"
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
	// Get user by email And get series
	GetUserByEmailAndSeries(ctx context.Context, email string) (entity.User, error)
	// Get user by email and password
	GetUserByEmailAndPassword(email, password string) error
	// Create new user
	CreateUser(email, password string) (entity.User, error)
	// Update user hash id
	UpdateUserHashId(ctx context.Context, email string, hashId string) error
	// Update user by email
	UpdateUserByEmail(email string) error
	// Delete user by email
	DeleteUserByEmail(email string) error
	// Generate Random Salt
	GenerateRandomSalt(saltLength uint32) ([]byte, error)
	// Generate Hash From Password
	GenerateFromPassword(password string) (string, error)
	// Compare Hash And Password
	CompareHashAndPassword(password string, encodedHash string) (bool, error)
	// Decode Hash
	DecodeHash(encodedHash string) (p *common.HashParams, salt, hash []byte, err error)
	// Update user password
	UpdateUserPassword(ctx context.Context, id uint, password string) error
	// Update user providers
	UpdateUserProviders(ctx context.Context, id uint, providerIds []uint) error
}
