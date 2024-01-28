package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewSeriesRepositoryImpl(DB *gorm.DB) repository.SeriesRepository {
	return &seriesRepositoryImpl{DB: DB}
}

type seriesRepositoryImpl struct {
	*gorm.DB
}

// Create Series
func (seriesRepository *seriesRepositoryImpl) CreateSeries(ctx context.Context, series entity.Series) (entity.Series, error) {
	result := seriesRepository.DB.WithContext(ctx).Create(series)
	if result.RowsAffected == 0 {
		return entity.Series{}, nil
	}
	return series, nil
}

func (seriesRepository *seriesRepositoryImpl) UpdateSeriesById(id uint, series entity.Series) (entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) DeleteSeriesById(ctx context.Context, id uint) error {
	result := seriesRepository.DB.WithContext(ctx).Delete(&entity.Series{}, id)
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesById(ctx context.Context, id uint) (entity.Series, error) {
	var seriesResult entity.Series
	result := seriesRepository.DB.WithContext(ctx).First(&seriesResult, id)
	if result.RowsAffected == 0 {
		return entity.Series{}, nil
	}
	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublishDayIdAndSeriesType(publishDayId uint, seriesType string) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublishDayId(publishDayId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByGenreId(genreId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPersonId(personId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublisherId(publisherId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesBySeriesType(seriesType string) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByTitle(title string) (entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByHashId(hashId string) (entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetAllSeries(ctx context.Context) ([]entity.Series, error) {
	var seriesResult []entity.Series
	result := seriesRepository.DB.WithContext(ctx).Find(&seriesResult)
	if result.RowsAffected == 0 {
		return []entity.Series{}, nil
	}
	return seriesResult, nil
}