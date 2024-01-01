package impl

import (
	"context"
	"errors"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewPublisherRepositoryImpl(DB *gorm.DB) repository.PublisherRepoistory {
	return &publisherRepositoryImpl{DB: DB}
}

type publisherRepositoryImpl struct {
	*gorm.DB
}

func (publisherRepository *publisherRepositoryImpl) GetPublisherById(ctx context.Context, id uint) (entity.Publisher, error) {
	//TODO implement me
	panic("implement me")
}

func (publisherRepository *publisherRepositoryImpl) GetPublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error) {
	//TODO implement me
	panic("implement me")
}

func (publisherRepository *publisherRepositoryImpl) GetPublisherByName(ctx context.Context, name string) (entity.Publisher, error) {
	//TODO implement me
	panic("implement me")
}

func (publisherRepository *publisherRepositoryImpl) CreatePublisher(ctx context.Context, name string) (entity.Publisher, error) {
	var publisherResult entity.Publisher
	publisherResult.Name = name
	result := publisherRepository.DB.WithContext(ctx).Create(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher already exist")
		return entity.Publisher{}, errors.New("publisher already exist")
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) UpdatePublisherHashId(ctx context.Context, publisher entity.Publisher, hashId string) error {
	var publisherResult entity.Publisher
	result := publisherRepository.DB.WithContext(ctx).Where("publisher.id = ?", publisher.Id).Find(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher not found")
	}
	publisherResult.HashId = hashId
	result = publisherRepository.DB.WithContext(ctx).Save(&publisherResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
}

func (publisherRepository *publisherRepositoryImpl) UpdatePublisher(ctx context.Context, publisher entity.Publisher) (entity.Publisher, error) {
	//TODO implement me
	panic("implement me")
}

func (publisherRepository *publisherRepositoryImpl) DeletePublisherById(ctx context.Context, id uint) (entity.Publisher, error) {
	//TODO implement me
	panic("implement me")
}

func (publisherRepository *publisherRepositoryImpl) DeletePublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error) {
	//TODO implement me
	panic("implement me")
}
