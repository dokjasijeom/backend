package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type UserRecordSeriesService interface {
	// Get user record series by id
	GetUserRecordSeriesById(ctx context.Context, id uint) (entity.UserRecordSeries, error)
	// Get user record series by user id
	GetUserRecordSeriesByUserId(ctx context.Context, userId uint) ([]entity.UserRecordSeries, error)
	// Get user record series by user id and series id
	GetUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) (entity.UserRecordSeries, error)
	// Get user record series by user id and user record series id
	GetUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint) (entity.UserRecordSeries, error)
	// Create user record series
	CreateUserRecordSeries(ctx context.Context, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error)
	// Update user record series by id
	UpdateUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error)
	// Delete user record series by id
	DeleteUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint) error
	// Delete user record series by user id and series id
	DeleteUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) error
	// Update User Record Series
	UpdateUserRecordSeries(ctx context.Context, userId, id uint, request entity.UserRecordSeries) (entity.UserRecordSeries, error)
	// Get User Complete Records
	GetUserCompleteRecords(ctx context.Context, userId uint) ([]entity.UserRecordSeries, error)
}
