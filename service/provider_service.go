package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type ProviderService interface {
	// Create New Provider
	CreateProvider(ctx context.Context, name string, displayName string, description string, homepageUrl string) (entity.Provider, error)
	// Get All Provider
	GetAllProvider(ctx context.Context) ([]entity.Provider, error)
	// Delete Provider
	DeleteProvider(ctx context.Context, providerId uint) error
	// Update Provider
	UpdateProvider(ctx context.Context, providerId uint, name string, displayName string, description string, homepageUrl string) error
	// Get Provider By Id
	GetProviderById(ctx context.Context, providerId uint) (entity.Provider, error)
	// Get Provider By Name
	GetProviderByName(ctx context.Context, name string) (entity.Provider, error)
}
