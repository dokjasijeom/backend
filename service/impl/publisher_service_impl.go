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

func (publisherService *publisherServiceImpl) CreatePublisher(ctx context.Context, name string) (entity.Publisher, error) {
	config := configuration.New()

	result, err := publisherService.PublisherRepoistory.CreatePublisher(ctx, name)
	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT_PUBLISHER")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(result.Id)})

	if e != "" {
		_ = publisherService.UpdatePublisherHashId(ctx, result, e)
		result.HashId = e
	}
	if err != nil {
		panic(err)
	}
	return result, nil
}
