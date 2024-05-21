package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewUserRecordSeriesEpisodeRepositoryImpl(DB *gorm.DB) repository.UserRecordSeriesEpisodeRepository {
	return &userRecordSeriesEpisodeRepositoryImpl{DB: DB}
}

type userRecordSeriesEpisodeRepositoryImpl struct {
	*gorm.DB
}

// Get user record series episode by id
func (userRecordSeriesEpisodeRepository *userRecordSeriesEpisodeRepositoryImpl) GetUserRecordSeriesEpisodeById(ctx context.Context, id uint) (entity.UserRecordSeriesEpisode, error) {
	var userRecordSeriesEpisode entity.UserRecordSeriesEpisode
	result := userRecordSeriesEpisodeRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&userRecordSeriesEpisode)
	return userRecordSeriesEpisode, result.Error
}

// Create user record series episode
func (userRecordSeriesEpisodeRepository *userRecordSeriesEpisodeRepositoryImpl) CreateUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error) {
	result := userRecordSeriesEpisodeRepository.DB.WithContext(ctx).Create(&userRecordSeriesEpisode)
	return userRecordSeriesEpisode, result.Error
}

// Create bulk user record series episode
func (userRecordSeriesEpisodeRepository *userRecordSeriesEpisodeRepositoryImpl) CreateBulkUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisodes []entity.UserRecordSeriesEpisode) ([]entity.UserRecordSeriesEpisode, error) {
	result := userRecordSeriesEpisodeRepository.DB.WithContext(ctx).Create(&userRecordSeriesEpisodes)
	return userRecordSeriesEpisodes, result.Error
}

// Update user record series episode by id
func (userRecordSeriesEpisodeRepository *userRecordSeriesEpisodeRepositoryImpl) UpdateUserRecordSeriesEpisodeById(ctx context.Context, id uint, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error) {
	result := userRecordSeriesEpisodeRepository.DB.WithContext(ctx).Where("id = ?", id).Updates(&userRecordSeriesEpisode)
	return userRecordSeriesEpisode, result.Error
}

// Delete user record series episode by id
func (userRecordSeriesEpisodeRepository *userRecordSeriesEpisodeRepositoryImpl) DeleteUserRecordSeriesEpisodeById(ctx context.Context, id uint) error {
	result := userRecordSeriesEpisodeRepository.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.UserRecordSeriesEpisode{})
	return result.Error
}

// Delete user record series episode by user record series id
func (userRecordSeriesEpisodeRepository *userRecordSeriesEpisodeRepositoryImpl) DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx context.Context, userRecordSeriesId uint) error {
	result := userRecordSeriesEpisodeRepository.DB.WithContext(ctx).Where("user_record_series_id = ?", userRecordSeriesId).Delete(&entity.UserRecordSeriesEpisode{})
	return result.Error
}
