package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type SeriesRepository interface {
	// Create Series
	CreateSeries(ctx context.Context, series entity.Series) (entity.Series, error)
	// Update Series by Id
	UpdateSeriesById(id uint, series entity.Series) (entity.Series, error)
	// Delete Series by Id
	DeleteSeriesById(id uint) error
	// Get Series by Id
	GetSeriesById(ctx context.Context, id uint) (entity.Series, error)
	// Get All Series
	GetAllSeries() ([]entity.Series, error)
	// Get Series by HashId
	GetSeriesByHashId(hashId string) (entity.Series, error)
	// Get Series by Title
	GetSeriesByTitle(title string) (entity.Series, error)
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
	// Get Series by PublishDayId and SeriesType
	GetSeriesByPublishDayIdAndSeriesType(publishDayId uint, seriesType string) ([]entity.Series, error)
}
