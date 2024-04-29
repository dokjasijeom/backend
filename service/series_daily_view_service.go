package service

import "context"

type SeriesDailyViewService interface {
	// Upsert SeriesDailyView
	UpsertSeriesDailyView(ctx context.Context, seriesId uint, currentDate string) error
}
