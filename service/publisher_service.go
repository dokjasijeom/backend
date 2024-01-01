package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PublisherService interface {
	// create new publisher
	CreatePublisher(ctx context.Context, name string) (entity.Publisher, error)
}
