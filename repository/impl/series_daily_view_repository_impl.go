package impl

import (
	"context"
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
	result := repository.DB.WithContext(ctx).Exec(`
		INSERT INTO series_daily_views (series_id, view_date, view_count, created_at, updated_at)
		VALUES (?, ?, 1, NOW(), NOW())
		ON DUPLICATE KEY UPDATE view_count = view_count + 1, updated_at = NOW()
	`, seriesId, currentDate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
