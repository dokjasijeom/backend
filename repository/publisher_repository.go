package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PublisherRepoistory interface {
	// get publisher by id
	GetPublisherById(ctx context.Context, id uint) (entity.Publisher, error)
	// get publisher by hash id
	GetPublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error)
	// get publisher by name
	GetPublisherByName(ctx context.Context, name string) (entity.Publisher, error)
	// create new publisher
	CreatePublisher(ctx context.Context, name string) (entity.Publisher, error)
	// update publisher
	UpdatePublisher(ctx context.Context, publisher entity.Publisher) (entity.Publisher, error)
	// update publisher's hashId
	UpdatePublisherHashId(ctx context.Context, publisher entity.Publisher, hashId string) error
	// delete publisher by id
	DeletePublisherById(ctx context.Context, id uint) (entity.Publisher, error)
	// delete publisher by hash id
	DeletePublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error)
}
