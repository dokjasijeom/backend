package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PublishDayService interface {
	// create new publish day
	CreatePublishDay(ctx context.Context, day string, displayDay string, displayOrder uint) (entity.PublishDay, error)
	// get all publish day
	GetAllPublishDay(ctx context.Context) ([]entity.PublishDay, error)
	// delete publish day by id
	DeletePublishDay(ctx context.Context, publishDayId uint) error
	// update publish day
	UpdatePublishDay(ctx context.Context, publishDayId uint, day string, displayDay string, displayOrder uint) error
	// get publish day by id
	GetPublishDayById(ctx context.Context, publishDayId uint) (entity.PublishDay, error)
}
