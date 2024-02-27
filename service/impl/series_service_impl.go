package impl

import (
	"context"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"github.com/speps/go-hashids/v2"
	"log"
)

func NewSeriesServiceImpl(seriesRepository *repository.SeriesRepository) service.SeriesService {
	return &seriesServiceImpl{SeriesRepository: *seriesRepository}
}

type seriesServiceImpl struct {
	repository.SeriesRepository
}

// Create Series
func (seriesService *seriesServiceImpl) CreateSeries(ctx context.Context, series model.SeriesModel) (entity.Series, error) {
	config := configuration.New()

	var seriesEntity = entity.Series{
		Title:      series.Title,
		Thumbnail:  series.Thumbnail,
		SeriesType: series.SeriesType,
	}

	if series.Description != "" {
		seriesEntity.Description = series.Description
	}
	if series.ISBN != "" {
		seriesEntity.ISBN = series.ISBN
	}
	if series.ECN != "" {
		seriesEntity.ECN = series.ECN
	}

	result, err := seriesService.SeriesRepository.CreateSeries(ctx, seriesEntity, series)
	if err != nil {
		exception.PanicLogging(err)
		return entity.Series{}, nil
	}

	log.Println("Series ID: ", result.Id)

	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT_SERIES")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(result.Id)})

	if e != "" {
		_ = seriesService.SeriesRepository.UpdateSeriesHashId(ctx, result.Id, e)
		result.HashId = e
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
	result, err := seriesService.SeriesRepository.GetSeriesById(ctx, id)
	if err != nil {
		return entity.Series{}, nil
	}
	return result, nil
}

// Get Series by HashId
func (seriesService *seriesServiceImpl) GetSeriesByHashId(ctx context.Context, hashId string) (entity.Series, error) {
	result, err := seriesService.SeriesRepository.GetSeriesByHashId(ctx, hashId)
	if err != nil {
		return entity.Series{}, nil
	}
	return result, nil

}

// Get all Series
func (seriesService *seriesServiceImpl) GetAllSeries(ctx context.Context) ([]entity.Series, error) {
	result, err := seriesService.SeriesRepository.GetAllSeries(ctx)
	if err != nil {
		return []entity.Series{}, nil
	}
	return result, nil
}

// Get PublishDay And SeriesType
func (seriesService *seriesServiceImpl) GetSeriesByPublishDayAndSeriesType(ctx context.Context, publishDay, seriesType string) ([]entity.Series, error) {
	result, err := seriesService.SeriesRepository.GetSeriesByPublishDayAndSeriesType(ctx, publishDay, seriesType)
	if err != nil {
		return []entity.Series{}, nil
	}
	return result, nil
}
