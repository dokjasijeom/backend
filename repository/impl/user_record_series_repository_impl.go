package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewUserRecordSeriesRepositoryImpl(DB *gorm.DB) repository.UserRecordSeriesRepository {
	return &userRecordSeriesRepositoryImpl{DB: DB}
}

type userRecordSeriesRepositoryImpl struct {
	*gorm.DB
}

// Get user record series by user id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesByUserId(ctx context.Context, userId uint) ([]entity.UserRecordSeries, error) {
	var userRecordSeries []entity.UserRecordSeries
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Get user record series by user id and series id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) (entity.UserRecordSeries, error) {
	var userRecordSeries entity.UserRecordSeries
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ? AND series_id = ?", userId, seriesId).First(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Create user record series
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) CreateUserRecordSeries(ctx context.Context, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Create(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Update user record series by id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) UpdateUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).Updates(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Delete user record series by id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) DeleteUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint) error {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).Delete(&entity.UserRecordSeries{})
	return result.Error
}

// Delete user record series by user id and series id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) DeleteUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) error {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ? AND series_id = ?", userId, seriesId).Delete(&entity.UserRecordSeries{})
	return result.Error
}
