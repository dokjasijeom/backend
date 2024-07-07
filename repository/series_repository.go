package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
)

type SeriesRepository interface {
	// Create Series
	CreateSeries(ctx context.Context, series entity.Series, model model.SeriesModel) (entity.Series, error)
	// Update Series HashId
	UpdateSeriesHashId(ctx context.Context, id uint, hashId string) error
	// Update Series by Id
	UpdateSeriesById(ctx context.Context, id uint, series entity.Series, model model.SeriesModel) (entity.Series, error)
	// Delete Series by Id
	DeleteSeriesById(ctx context.Context, id uint) error
	// Get Series by Id
	GetSeriesById(ctx context.Context, id uint) (entity.Series, error)
	// Get Backoffice All Series
	GetBackofficeAllSeries(ctx context.Context) ([]entity.Series, error)
	// Get All Series
	GetAllSeries(ctx context.Context, page, pageSize int) (model.SeriesWithPagination, error)
	// Get All Category Series
	GetAllCategorySeries(ctx context.Context, seriesType entity.SeriesType, sort, genre string, providers []string, page, pageSize int) (model.SeriesWithPagination, error)
	// Get Series by HashId
	GetSeriesByHashId(ctx context.Context, hashId string) (entity.Series, error)
	// Get Series by Title
	GetSeriesByTitle(ctx context.Context, title string) ([]entity.Series, error)
	// Get Series Id And Titles by Title
	GetSeriesIdAndTitlesByTitle(ctx context.Context, title string) ([]entity.Series, error)
	// Get Series by SeriesType
	GetSeriesBySeriesType(seriesType string) ([]entity.Series, error)
	// Get Series by PublisherId
	GetSeriesByPublisherId(publisherId uint) ([]entity.Series, error)
	// Get Series by PersonId
	GetSeriesByPersonId(personId uint) ([]entity.Series, error)
	// Get Series by GenreId
	GetSeriesByGenreId(genreId uint) ([]entity.Series, error)
	// Get Series by PublishDayId
	GetSeriesByPublishDayId(publishDayId uint) ([]entity.Series, error)
	// Get Series by PublishDays And SeriesType
	GetSeriesByPublishDayAndSeriesType(ctx context.Context, publishDay, seriesType string, page, pageSize int) (model.SeriesWithPagination, error)
	// Get New Episode Update Provider Series
	GetNewEpisodeUpdateProviderSeries(ctx context.Context, provider, seriesType string, page, pageSize int) (model.SeriesWithPagination, error)
	// Has Like Series
	HasLikeSeries(ctx context.Context, userId uint, seriesId uint) (bool, error)
	// Like Series
	LikeSeries(ctx context.Context, userId uint, seriesId uint) error
	// Unlike Series
	UnlikeSeries(ctx context.Context, userId uint, seriesId uint) error
}
