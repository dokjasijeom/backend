package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type EpisodeService interface {
	// create new episode
	CreateEpisode(ctx context.Context, episodeNumber uint) (entity.Episode, error)
	// get all episode
	GetAllEpisode(ctx context.Context) ([]entity.Episode, error)
	// get episodes by series id
	GetEpisodesBySeriesId(ctx context.Context, seriesId uint) ([]entity.Episode, error)
	// delete episode by id
	DeleteEpisode(ctx context.Context, episodeId uint) error
	// update episode
	UpdateEpisode(ctx context.Context, episodeId uint, episodeNumber uint) error
	// get episode by id
	GetEpisodeById(ctx context.Context, episodeId uint) (entity.Episode, error)
	// create bulk new episode
	CreateBulkEpisode(ctx context.Context, seriesId uint, toEpisodeNumber uint) ([]entity.Episode, error)
}
