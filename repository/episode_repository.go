package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type EpisodeRepository interface {
	// Create New Episode
	CreateEpisode(ctx context.Context, episodeNumber uint) (entity.Episode, error)
	// Get All Episode
	GetAllEpisode(ctx context.Context) ([]entity.Episode, error)
	// Get Episodes By Series Id
	GetEpisodesBySeriesId(ctx context.Context, seriesId uint) ([]entity.Episode, error)
	// Delete Episode
	DeleteEpisode(ctx context.Context, episodeId uint) error
	// Update Episode
	UpdateEpisode(ctx context.Context, episodeId uint, episodeNumber uint) error
	// Get Episode By Id
	GetEpisodeById(ctx context.Context, episodeId uint) (entity.Episode, error)
	// Create Bulk New Episode
	CreateBulkEpisode(ctx context.Context, seriesId, toEpisodeNumber uint) ([]entity.Episode, error)
}
