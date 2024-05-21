package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewEpisodeServiceImpl(episodeRepository *repository.EpisodeRepository) service.EpisodeService {
	return &episodeServiceImpl{EpisodeRepository: *episodeRepository}
}

type episodeServiceImpl struct {
	repository.EpisodeRepository
}

func (episodeService *episodeServiceImpl) CreateEpisode(ctx context.Context, episodeNumber uint) (entity.Episode, error) {
	return episodeService.EpisodeRepository.CreateEpisode(ctx, episodeNumber)
}

func (episodeService *episodeServiceImpl) GetAllEpisode(ctx context.Context) ([]entity.Episode, error) {
	return episodeService.EpisodeRepository.GetAllEpisode(ctx)
}

func (episodeService *episodeServiceImpl) GetEpisodesBySeriesId(ctx context.Context, seriesId uint) ([]entity.Episode, error) {
	return episodeService.EpisodeRepository.GetEpisodesBySeriesId(ctx, seriesId)
}

func (episodeService *episodeServiceImpl) DeleteEpisode(ctx context.Context, episodeId uint) error {
	return episodeService.EpisodeRepository.DeleteEpisode(ctx, episodeId)
}

func (episodeService *episodeServiceImpl) UpdateEpisode(ctx context.Context, episodeId uint, episodeNumber uint) error {
	return episodeService.EpisodeRepository.UpdateEpisode(ctx, episodeId, episodeNumber)
}

func (episodeService *episodeServiceImpl) GetEpisodeById(ctx context.Context, episodeId uint) (entity.Episode, error) {
	return episodeService.EpisodeRepository.GetEpisodeById(ctx, episodeId)
}

func (episodeService *episodeServiceImpl) CreateBulkEpisode(ctx context.Context, seriesId, toEpisodeNumber uint) ([]entity.Episode, error) {
	return episodeService.EpisodeRepository.CreateBulkEpisode(ctx, seriesId, toEpisodeNumber)
}
