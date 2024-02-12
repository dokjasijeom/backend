package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewPublishDayRepositoryImpl(DB *gorm.DB) repository.PublishDayRepository {
	return &publishDayRepositoryImpl{DB: DB}
}

type publishDayRepositoryImpl struct {
	*gorm.DB
}

func (publishDayRepository *publishDayRepositoryImpl) CreatePublishDay(ctx context.Context, day string, displayDay string, displayOrder uint) (entity.PublishDay, error) {
	publishDay := entity.PublishDay{
		Day:          day,
		DisplayDay:   displayDay,
		DisplayOrder: displayOrder,
	}
	result := publishDayRepository.DB.WithContext(ctx).Create(&publishDay)
	if result.Error != nil {
		return entity.PublishDay{}, result.Error
	}
	return publishDay, nil
}

func (publishDayRepository *publishDayRepositoryImpl) GetAllPublishDay(ctx context.Context) ([]entity.PublishDay, error) {
	var publishDayResult []entity.PublishDay
	result := publishDayRepository.DB.WithContext(ctx).Find(&publishDayResult)
	if result.Error != nil {
		return nil, result.Error
	}
	return publishDayResult, nil
}

func (publishDayRepository *publishDayRepositoryImpl) DeletePublishDay(ctx context.Context, publishDayId uint) error {
	result := publishDayRepository.DB.WithContext(ctx).Where("publish_day.id = ?", publishDayId).Delete(&entity.PublishDay{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (publishDayRepository *publishDayRepositoryImpl) UpdatePublishDay(ctx context.Context, publishDayId uint, day string, displayDay string, displayOrder uint) error {
	result := publishDayRepository.DB.WithContext(ctx).Model(&entity.PublishDay{}).Where("publish_day.id = ?", publishDayId).Updates(entity.PublishDay{
		Day:          day,
		DisplayDay:   displayDay,
		DisplayOrder: displayOrder,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (publishDayRepository *publishDayRepositoryImpl) GetPublishDayById(ctx context.Context, publishDayId uint) (entity.PublishDay, error) {
	var publishDayResult entity.PublishDay
	result := publishDayRepository.DB.WithContext(ctx).Where("publish_day.id = ?", publishDayId).Find(&publishDayResult)
	if result.Error != nil {
		return entity.PublishDay{}, result.Error
	}
	return publishDayResult, nil
}
