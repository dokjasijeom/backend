package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type UserRecordSeriesEpisodeService interface {
	// Get user record series episode by id
	GetUserRecordSeriesEpisodeById(ctx context.Context, id uint) (entity.UserRecordSeriesEpisode, error)
	// Create user record series episode
	CreateUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error)
	// Create bulk user record series episode
	CreateBulkUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisodes []entity.UserRecordSeriesEpisode) ([]entity.UserRecordSeriesEpisode, error)
	// Update user record series episode by id
	UpdateUserRecordSeriesEpisodeById(ctx context.Context, id uint, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error)
	// Delete user record series episode by id
	DeleteUserRecordSeriesEpisodeById(ctx context.Context, id uint) error
	// Delete user record series episode by user record series id
	DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx context.Context, userRecordSeriesId uint) error
}
