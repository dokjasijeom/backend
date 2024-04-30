package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewSeriesDailyViewRepositoryImpl(DB *gorm.DB) repository.SeriesDailyViewRepository {
	return &seriesDailyViewRepositoryImpl{DB: DB}
}

type seriesDailyViewRepositoryImpl struct {
	DB *gorm.DB
}

// Upsert SeriesDailyView
func (repository *seriesDailyViewRepositoryImpl) UpsertSeriesDailyView(ctx context.Context, seriesId uint, currentDate string) error {
	// if not exist series_id and view_date, insert new record
	// if exist series_id and view_date, increase view_count
	var seriesDailyView entity.SeriesDailyView
	seriesDailyView.SeriesId = seriesId
	seriesDailyView.ViewDate = currentDate
	seriesDailyView.ViewCount = 0
	result := repository.DB.WithContext(ctx).Model(&entity.SeriesDailyView{}).Where("series_id = ? AND view_date = ?", seriesId, currentDate).FirstOrCreate(&seriesDailyView)
	if result.Error != nil {
		return result.Error
	}

	seriesDailyView.ViewCount++
	result = repository.DB.WithContext(ctx).Save(&seriesDailyView)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
