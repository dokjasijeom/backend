package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewSeriesServiceImpl(seriesRepository *repository.SeriesRepository) service.SeriesService {
	return &seriesServiceImpl{SeriesRepository: *seriesRepository}
}

type seriesServiceImpl struct {
	repository.SeriesRepository
}

// Create Series
func (seriesService *seriesServiceImpl) CreateSeries(ctx context.Context, series model.SeriesModel) (entity.Series, error) {
	var seriesEntity = entity.Series{
		Title:       series.Title,
		Description: series.Description,
		Thumbnail:   series.Thumbnail,
		ISBN:        series.ISBN,
		ECNNumber:   series.ECNNumber,
		SeriesType:  series.SeriesType,
	}

	result, err := seriesService.SeriesRepository.CreateSeries(seriesEntity)
	if err != nil {
		return entity.Series{}, nil
	}
	return result, nil
}

// Update Series by Id
func (seriesService *seriesServiceImpl) UpdateSeriesById(ctx context.Context, id uint, series entity.Series) (entity.Series, error) {
	panic("implement me")
}

// Delete Series by Id
func (seriesService *seriesServiceImpl) DeleteSeriesById(ctx context.Context, id uint) error {
	panic("implement me")
}

// Get Series by Id
func (seriesService *seriesServiceImpl) GetSeriesById(ctx context.Context, id uint) (entity.Series, error) {
	result, err := seriesService.SeriesRepository.GetSeriesById(id)
	if err != nil {
		return entity.Series{}, nil
	}
	return result, nil
}

// Get all Series
func (seriesService *seriesServiceImpl) GetAllSeries(ctx context.Context) ([]entity.Series, error) {
	result, _ := seriesService.SeriesRepository.GetAllSeries()
	if result == nil {
		return []entity.Series{}, nil
	}
	return result, nil
}
