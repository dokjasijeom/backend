package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewProviderRepositoryImpl(DB *gorm.DB) repository.ProviderRepository {
	return &providerRepositoryImpl{DB: DB}
}

type providerRepositoryImpl struct {
	*gorm.DB
}

func (providerRepository *providerRepositoryImpl) CreateProvider(ctx context.Context, name string, displayName string, description string, homepageUrl string) (entity.Provider, error) {
	var providerResult entity.Provider
	providerResult.Name = name
	providerResult.DisplayName = displayName
	providerResult.Description = description
	providerResult.HomepageUrl = homepageUrl

	result := providerRepository.DB.WithContext(ctx).Create(&providerResult)
	if result.Error != nil {
		providerResult.Id = 0
		return providerResult, result.Error
	}
	return providerResult, nil
}

func (providerRepository *providerRepositoryImpl) GetAllProvider(ctx context.Context) ([]entity.Provider, error) {
	var providerResult []entity.Provider
	result := providerRepository.DB.WithContext(ctx).Find(&providerResult)

	if result.Error != nil {
		return nil, result.Error
	}

	return providerResult, nil
}

func (providerRepository *providerRepositoryImpl) DeleteProvider(ctx context.Context, providerId uint) error {
	result := providerRepository.DB.WithContext(ctx).Delete(&entity.Provider{}, providerId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (providerRepository *providerRepositoryImpl) UpdateProvider(ctx context.Context, providerId uint, name string, displayName string, description string, homepageUrl string) error {
	result := providerRepository.DB.WithContext(ctx).Model(&entity.Provider{}).Where("id = ?", providerId).Updates(map[string]interface{}{"name": name, "display_name": displayName, "description": description, "homepage_url": homepageUrl})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (providerRepository *providerRepositoryImpl) GetProviderById(ctx context.Context, providerId uint) (entity.Provider, error) {
	var providerResult entity.Provider
	result := providerRepository.DB.WithContext(ctx).Where("id = ?", providerId).Find(&providerResult)

	if result.Error != nil {
		return entity.Provider{}, result.Error
	}

	return providerResult, nil
}

func (providerRepository *providerRepositoryImpl) GetProviderByName(ctx context.Context, name string) (entity.Provider, error) {
	var providerResult entity.Provider
	result := providerRepository.DB.WithContext(ctx).Where("name = ?", name).Find(&providerResult)

	if result.Error != nil {
		return entity.Provider{}, result.Error
	}

	return providerResult, nil
}

func (providerRepository *providerRepositoryImpl) UpdateHashIdProvider(ctx context.Context, providerId uint, hashId string) error {
	result := providerRepository.DB.WithContext(ctx).Model(&entity.Provider{}).Where("id = ?", providerId).Update("hash_id", hashId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
