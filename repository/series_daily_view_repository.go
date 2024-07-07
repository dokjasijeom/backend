package repository

import (
	"context"
	"time"
)

type SeriesDailyViewRepository interface {
	// Upsert SeriesDailyView
	UpsertSeriesDailyView(ctx context.Context, seriesId uint, currentDate time.Time) error
}
