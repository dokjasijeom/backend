package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewUserRecordSeriesServiceImpl(userRecordSeriesRepository *repository.UserRecordSeriesRepository) service.UserRecordSeriesService {
	return &userRecordSeriesServiceImpl{UserRecordSeriesRepository: *userRecordSeriesRepository}
}

type userRecordSeriesServiceImpl struct {
	repository.UserRecordSeriesRepository
}

// Get user record series by user id
func (userRecordSeriesService *userRecordSeriesServiceImpl) GetUserRecordSeriesByUserId(ctx context.Context, userId uint) ([]entity.UserRecordSeries, error) {
	return userRecordSeriesService.UserRecordSeriesRepository.GetUserRecordSeriesByUserId(ctx, userId)
}

// Get user record series by user id and series id
func (userRecordSeriesService *userRecordSeriesServiceImpl) GetUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) (entity.UserRecordSeries, error) {
	return userRecordSeriesService.UserRecordSeriesRepository.GetUserRecordSeriesByUserIdAndSeriesId(ctx, userId, seriesId)
}

// Create user record series
func (userRecordSeriesService *userRecordSeriesServiceImpl) CreateUserRecordSeries(ctx context.Context, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	return userRecordSeriesService.UserRecordSeriesRepository.CreateUserRecordSeries(ctx, userRecordSeries)
}

// Update user record series by id
func (userRecordSeriesService *userRecordSeriesServiceImpl) UpdateUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	return userRecordSeriesService.UserRecordSeriesRepository.UpdateUserRecordSeriesByUserIdAndId(ctx, userId, id, userRecordSeries)
}

// Delete user record series by id
func (userRecordSeriesService *userRecordSeriesServiceImpl) DeleteUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint) error {
	return userRecordSeriesService.UserRecordSeriesRepository.DeleteUserRecordSeriesByUserIdAndId(ctx, userId, id)
}

// Delete user record series by user id and series id
func (userRecordSeriesService *userRecordSeriesServiceImpl) DeleteUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) error {
	return userRecordSeriesService.UserRecordSeriesRepository.DeleteUserRecordSeriesByUserIdAndSeriesId(ctx, userId, seriesId)
}
