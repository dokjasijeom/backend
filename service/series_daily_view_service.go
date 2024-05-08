package service

import (
	"context"
	"time"
)

type SeriesDailyViewService interface {
	// Upsert SeriesDailyView
	UpsertSeriesDailyView(ctx context.Context, seriesId uint, currentDate time.Time) error
}
