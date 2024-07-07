package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewPublishDayServiceImpl(publishDayRepository *repository.PublishDayRepository) service.PublishDayService {
	return &publishDayServiceImpl{PublishDayRepository: *publishDayRepository}
}

type publishDayServiceImpl struct {
	repository.PublishDayRepository
}

func (publishDayService *publishDayServiceImpl) CreatePublishDay(ctx context.Context, day string, displayDay string, displayOrder uint) (entity.PublishDay, error) {
	return publishDayService.PublishDayRepository.CreatePublishDay(ctx, day, displayDay, displayOrder)
}

func (publishDayService *publishDayServiceImpl) GetAllPublishDay(ctx context.Context) ([]entity.PublishDay, error) {
	return publishDayService.PublishDayRepository.GetAllPublishDay(ctx)
}

func (publishDayService *publishDayServiceImpl) DeletePublishDay(ctx context.Context, publishDayId uint) error {
	return publishDayService.PublishDayRepository.DeletePublishDay(ctx, publishDayId)
}

func (publishDayService *publishDayServiceImpl) UpdatePublishDay(ctx context.Context, publishDayId uint, day string, displayDay string, displayOrder uint) error {
	return publishDayService.PublishDayRepository.UpdatePublishDay(ctx, publishDayId, day, displayDay, displayOrder)
}

func (publishDayService *publishDayServiceImpl) GetPublishDayById(ctx context.Context, publishDayId uint) (entity.PublishDay, error) {
	return publishDayService.PublishDayRepository.GetPublishDayById(ctx, publishDayId)
}
