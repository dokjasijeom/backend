package repository

import "context"

type SeriesDailyViewRepository interface {
	// Upsert SeriesDailyView
	UpsertSeriesDailyView(ctx context.Context, seriesId uint, currentDate string) error
}
