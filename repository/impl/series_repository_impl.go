package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
	"log"
)

func NewSeriesRepositoryImpl(DB *gorm.DB) repository.SeriesRepository {
	return &seriesRepositoryImpl{DB: DB}
}

type seriesRepositoryImpl struct {
	*gorm.DB
}

// Create Series
func (seriesRepository *seriesRepositoryImpl) CreateSeries(ctx context.Context, series entity.Series, model model.SeriesModel) (entity.Series, error) {
	var seriesResult entity.Series
	seriesResult = series

	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Create(&seriesResult)

	if model.GenreId != 0 {
		err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Genres").Append(&entity.Genre{Id: model.GenreId})
		if err != nil {
			log.Println("장르 연결 실패")
			exception.PanicLogging(err)
			//return entity.Series{}, err
		}
	}

	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return entity.Series{}, result.Error
	}
	return seriesResult, nil
}

// Update series hash id
func (seriesRepository *seriesRepositoryImpl) UpdateSeriesHashId(ctx context.Context, id uint, hashId string) error {
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("id = ?", id).Updates(map[string]interface{}{"hash_id": hashId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (seriesRepository *seriesRepositoryImpl) UpdateSeriesById(id uint, series entity.Series) (entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) DeleteSeriesById(ctx context.Context, id uint) error {
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Delete(&entity.Series{}, id)
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesById(ctx context.Context, id uint) (entity.Series, error) {
	var seriesResult entity.Series
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).First(&seriesResult, id)
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
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Find(&seriesResult)
	if result.RowsAffected == 0 {
		return []entity.Series{}, nil
	}
	return seriesResult, nil
}
