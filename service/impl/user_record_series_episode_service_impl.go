package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewUserRecordSeriesEpisodeServiceImpl(userRecordSeriesEpisodeRepository *repository.UserRecordSeriesEpisodeRepository) service.UserRecordSeriesEpisodeService {
	return &userRecordSeriesEpisodeServiceImpl{UserRecordSeriesEpisodeRepository: *userRecordSeriesEpisodeRepository}
}

type userRecordSeriesEpisodeServiceImpl struct {
	repository.UserRecordSeriesEpisodeRepository
}

// Get user record series episode by id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) GetUserRecordSeriesEpisodeById(ctx context.Context, id uint) (entity.UserRecordSeriesEpisode, error) {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.GetUserRecordSeriesEpisodeById(ctx, id)
}

// Create user record series episode
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) CreateUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error) {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.CreateUserRecordSeriesEpisode(ctx, userRecordSeriesEpisode)
}

// Create bulk user record series episode
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) CreateBulkUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisodes []entity.UserRecordSeriesEpisode) ([]entity.UserRecordSeriesEpisode, error) {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.CreateBulkUserRecordSeriesEpisode(ctx, userRecordSeriesEpisodes)
}

// Update user record series episode by id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) UpdateUserRecordSeriesEpisodeById(ctx context.Context, id uint, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error) {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.UpdateUserRecordSeriesEpisodeById(ctx, id, userRecordSeriesEpisode)
}

// Delete user record series episode by id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) DeleteUserRecordSeriesEpisodeById(ctx context.Context, id uint) error {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.DeleteUserRecordSeriesEpisodeById(ctx, id)
}

// Delete user record series episode by user record series id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx context.Context, userRecordSeriesId uint) error {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx, userRecordSeriesId)
}
