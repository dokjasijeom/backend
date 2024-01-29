package impl

import (
	"context"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"github.com/speps/go-hashids/v2"
)

func NewProviderServiceImpl(providerRepository *repository.ProviderRepository) service.ProviderService {
	return &providerServiceImpl{ProviderRepository: *providerRepository}
}

type providerServiceImpl struct {
	repository.ProviderRepository
}

func (providerService *providerServiceImpl) CreateProvider(ctx context.Context, name string, displayName string, description string, homepageUrl string) (entity.Provider, error) {
	config := configuration.New()

	provider, err := providerService.ProviderRepository.CreateProvider(ctx, name, displayName, description, homepageUrl)
	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT_PROVIDER")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(provider.Id)})

	if e != "" {
		_ = providerService.ProviderRepository.UpdateHashIdProvider(ctx, provider.Id, e)
	}
	if err != nil {
		exception.PanicLogging(err)
	}
	if provider.Id == 0 {
		exception.PanicLogging("provider is empty")
	}

	return provider, nil
}

func (providerService *providerServiceImpl) GetAllProvider(ctx context.Context) ([]entity.Provider, error) {
	result, err := providerService.ProviderRepository.GetAllProvider(ctx)
	if err != nil {
		exception.PanicLogging(err)
		return nil, err
	}
	return result, nil
}

func (providerService *providerServiceImpl) DeleteProvider(ctx context.Context, providerId uint) error {
	err := providerService.ProviderRepository.DeleteProvider(ctx, providerId)
	if err != nil {
		exception.PanicLogging(err)
		return err
	}
	return nil
}

func (providerService *providerServiceImpl) UpdateProvider(ctx context.Context, providerId uint, name string, displayName string, description string, homepageUrl string) error {
	err := providerService.ProviderRepository.UpdateProvider(ctx, providerId, name, displayName, description, homepageUrl)
	if err != nil {
		exception.PanicLogging(err)
		return err
	}
	return nil
}

func (providerService *providerServiceImpl) GetProviderById(ctx context.Context, providerId uint) (entity.Provider, error) {
	result, err := providerService.ProviderRepository.GetProviderById(ctx, providerId)
	if err != nil {
		exception.PanicLogging(err)
		return entity.Provider{}, err
	}
	return result, nil
}

func (providerService *providerServiceImpl) GetProviderByName(ctx context.Context, name string) (entity.Provider, error) {
	result, err := providerService.ProviderRepository.GetProviderByName(ctx, name)
	if err != nil {
		exception.PanicLogging(err)
		return entity.Provider{}, err
	}
	return result, nil
}
