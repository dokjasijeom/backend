package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewUserProfileRepositoryImpl(DB *gorm.DB) repository.UserProfileRepository {
	return &userProfileRepositoryImpl{DB: DB}
}

type userProfileRepositoryImpl struct {
	DB *gorm.DB
}

func (upr *userProfileRepositoryImpl) UpdateUserProfile(ctx context.Context, id uint, request model.UserUpdateRequestModel) error {
	// Update user profile
	var userProfile entity.UserProfile
	upr.DB.WithContext(ctx).Where("user_id = ?", id).First(&userProfile)
	if userProfile.UserId == 0 {
		userProfile.UserId = id
	}
	userProfile.Username = request.Username
	if request.Avatar != "" {
		userProfile.Avatar = request.Avatar
	}
	result := upr.DB.Save(&userProfile)

	return result.Error
}
