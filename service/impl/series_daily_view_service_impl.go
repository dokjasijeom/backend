package impl

import (
	"context"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"time"
)

func NewSeriesDailyViewServiceImpl(seriesDailyViewRepository *repository.SeriesDailyViewRepository) service.SeriesDailyViewService {
	return &seriesDailyViewServiceImpl{SeriesDailyViewRepository: *seriesDailyViewRepository}
}

type seriesDailyViewServiceImpl struct {
	repository.SeriesDailyViewRepository
}

func (service *seriesDailyViewServiceImpl) UpsertSeriesDailyView(ctx context.Context, seriesId uint, currentDate time.Time) error {
	err := service.SeriesDailyViewRepository.UpsertSeriesDailyView(ctx, seriesId, currentDate)
	if err != nil {
		panic(err)
	}
	return nil
}
