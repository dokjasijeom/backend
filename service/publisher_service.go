package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PublisherService interface {
	// create new publisher
	CreatePublisher(ctx context.Context, name, description, homepageUrl string) (entity.Publisher, error)
	// get publisher by id
	GetPublisherById(ctx context.Context, id uint) (entity.Publisher, error)
	// get publisher by hash id
	GetPublisherByHashId(ctx context.Context, hashId string) (entity.Publisher, error)
	// get publisher by name
	GetPublisherByName(ctx context.Context, name string) (entity.Publisher, error)
	// get all publisher
	GetAllPublisher(ctx context.Context) ([]entity.Publisher, error)
	// update publisher
	UpdatePublisher(ctx context.Context, id uint, name, description, homepageUrl string) (entity.Publisher, error)
	// delete publisher by id
	DeletePublisherById(ctx context.Context, id uint) error
	// delete publisher by hash id
	DeletePublisherByHashId(ctx context.Context, hashId string) error
}
