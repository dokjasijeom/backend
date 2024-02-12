package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PublishDayRepository interface {
	// Create New PublishDay
	CreatePublishDay(ctx context.Context, day string, displayDay string, displayOrder uint) (entity.PublishDay, error)
	// Get All PublishDay
	GetAllPublishDay(ctx context.Context) ([]entity.PublishDay, error)
	// Delete PublishDay
	DeletePublishDay(ctx context.Context, publishDayId uint) error
	// Update PublishDay
	UpdatePublishDay(ctx context.Context, publishDayId uint, day string, displayDay string, displayOrder uint) error
	// Get PublishDay By Id
	GetPublishDayById(ctx context.Context, publishDayId uint) (entity.PublishDay, error)
}
