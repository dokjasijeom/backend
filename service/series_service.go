package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
)

type SeriesService interface {
	// Create Series
	CreateSeries(ctx context.Context, series model.SeriesModel) (entity.Series, error)
	// Update Series by Id
	UpdateSeriesById(ctx context.Context, id uint, series entity.Series, model model.SeriesModel) (entity.Series, error)
	// Delete Series by Id
	DeleteSeriesById(ctx context.Context, id uint) error
	// Get Series by Id
	GetSeriesById(ctx context.Context, id uint) (entity.Series, error)
	// Get Series by HashId
	GetSeriesByHashId(ctx context.Context, hashId string) (entity.Series, error)
	// Get all Series
	GetAllSeries(ctx context.Context) ([]entity.Series, error)
	// Get PublishDay And SeriesType
	GetSeriesByPublishDayAndSeriesType(ctx context.Context, publishDay, seriesType string) ([]entity.Series, error)
	// Get New Episode Update Provider Series
	GetNewEpisodeUpdateProviderSeries(ctx context.Context, provider, seriesType string) ([]entity.Series, error)
	// Like Series
	LikeSeries(ctx context.Context, userId uint, seriesId uint) error
	// Get Series By Title
	GetSeriesByTitle(ctx context.Context, title string) ([]entity.Series, error)
	// Get Series Id And Title by Title
	GetSeriesIdAndTitlesByTitle(ctx context.Context, title string) ([]entity.Series, error)
}
