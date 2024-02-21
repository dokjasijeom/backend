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
	var publisherResult entity.Publisher
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Where("id = ?", id).Find(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher not found")
		return entity.Publisher{}, errors.New("publisher not found")
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) GetPublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error) {
	var publisherResult entity.Publisher
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Where("hash_id = ?", hashId).Find(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher not found")
		return entity.Publisher{}, errors.New("publisher not found")
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) GetPublisherByName(ctx context.Context, name string) (entity.Publisher, error) {
	var publisherResult entity.Publisher
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Where("name = ?", name).Find(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher not found")
		return entity.Publisher{}, errors.New("publisher not found")
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) GetAllPublisher(ctx context.Context) ([]entity.Publisher, error) {
	var publisherResult []entity.Publisher
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Find(&publisherResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return nil, result.Error
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) CreatePublisher(ctx context.Context, name, description, homepageUrl string) (entity.Publisher, error) {
	var publisherResult entity.Publisher
	publisherResult.Name = name
	publisherResult.Description = description
	publisherResult.HomepageUrl = homepageUrl

	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Create(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher already exist")
		return entity.Publisher{}, errors.New("publisher already exist")
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) UpdatePublisherHashId(ctx context.Context, id uint, hashId string) error {
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Where("id = ?", id).Updates(map[string]interface{}{"hash_id": hashId})
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
}

func (publisherRepository *publisherRepositoryImpl) UpdatePublisher(ctx context.Context, id uint, name, description, homepageUrl string) (entity.Publisher, error) {
	var publisherResult entity.Publisher
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Where("id = ?", id).Find(&publisherResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("publisher not found")
	}
	publisherResult.Name = name
	publisherResult.Description = description
	publisherResult.HomepageUrl = homepageUrl
	result = publisherRepository.DB.WithContext(ctx).Save(&publisherResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return entity.Publisher{}, result.Error
	}
	return publisherResult, nil
}

func (publisherRepository *publisherRepositoryImpl) DeletePublisherById(ctx context.Context, id uint) error {
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Delete(&entity.Publisher{}, id)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
}

func (publisherRepository *publisherRepositoryImpl) DeletePublisherByHashId(ctx context.Context, hashId string) error {
	result := publisherRepository.DB.WithContext(ctx).Model(&entity.Publisher{}).Where("hash_id = ?", hashId).Delete(&entity.Publisher{})
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}
	return nil
}
