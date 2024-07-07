package impl

import (
	"context"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"github.com/speps/go-hashids/v2"
)

func NewPublisherServiceImpl(publisherRepoistory *repository.PublisherRepoistory) service.PublisherService {
	return &publisherServiceImpl{PublisherRepoistory: *publisherRepoistory}
}

type publisherServiceImpl struct {
	repository.PublisherRepoistory
}

func (publisherService *publisherServiceImpl) CreatePublisher(ctx context.Context, name, description, homepageUrl string) (entity.Publisher, error) {
	config := configuration.New()

	result, err := publisherService.PublisherRepoistory.CreatePublisher(ctx, name, description, homepageUrl)
	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT_PUBLISHER")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(result.Id)})

	if e != "" {
		_ = publisherService.PublisherRepoistory.UpdatePublisherHashId(ctx, result.Id, e)
		result.HashId = e
	}
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (publisherService *publisherServiceImpl) GetPublisherById(ctx context.Context, id uint) (entity.Publisher, error) {
	return publisherService.PublisherRepoistory.GetPublisherById(ctx, id)
}

func (publisherService *publisherServiceImpl) GetPublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error) {
	return publisherService.PublisherRepoistory.GetPublisherByHashId(ctx, hashId)
}

func (publisherService *publisherServiceImpl) GetPublisherByName(ctx context.Context, name string) (entity.Publisher, error) {
	return publisherService.PublisherRepoistory.GetPublisherByName(ctx, name)
}

func (publisherService *publisherServiceImpl) GetAllPublisher(ctx context.Context) ([]entity.Publisher, error) {
	return publisherService.PublisherRepoistory.GetAllPublisher(ctx)
}

func (publisherService *publisherServiceImpl) UpdatePublisher(ctx context.Context, id uint, name, description, homepageUrl string) (entity.Publisher, error) {
	return publisherService.PublisherRepoistory.UpdatePublisher(ctx, id, name, description, homepageUrl)
}

func (publisherService *publisherServiceImpl) DeletePublisherById(ctx context.Context, id uint) error {
	return publisherService.PublisherRepoistory.DeletePublisherById(ctx, id)
}

func (publisherService *publisherServiceImpl) DeletePublisherByHashId(ctx context.Context, hashId string) error {
	return publisherService.PublisherRepoistory.DeletePublisherByHashId(ctx, hashId)
}
