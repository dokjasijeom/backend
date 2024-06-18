package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
)

type UserService interface {
	// User Authentication
	AuthenticateUser(ctx context.Context, email, password string) (entity.User, error)
	// User Create
	CreateUser(ctx context.Context, email, password, comparePassword string) (entity.User, error)
	// Get User By Email
	GetUserByEmail(ctx context.Context, email string) entity.User
	// Get User By Email And Series
	GetUserByEmailAndSeries(ctx context.Context, email string) entity.User
	// Update User Profile
	UpdateUserProfile(ctx context.Context, id uint, request model.UserUpdateRequestModel) (bool, error)
	// Update User Password
	UpdateUserPassword(ctx context.Context, id uint, request model.UserUpdateRequestModel) (bool, error)
	// Update User Providers
	UpdateUserProviders(ctx context.Context, id uint, providerIds []uint) (bool, error)
	// Delete User Avatar
	DeleteUserAvatar(ctx context.Context, id uint) (bool, error)
}
