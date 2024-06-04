package repository

import (
	"context"
	"github.com/dokjasijeom/backend/model"
)

type UserProfileRepository interface {
	// Update User Profile
	UpdateUserProfile(ctx context.Context, id uint, request model.UserUpdateRequestModel) error
}
