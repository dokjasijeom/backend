package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type ProviderRepository interface {
	// Create New Provider
	CreateProvider(ctx context.Context, name string, displayName string, description string, homepageUrl string) (entity.Provider, error)
	// Get All Provider
	GetAllProvider(ctx context.Context) ([]entity.Provider, error)
	// Delete Provider
	DeleteProvider(ctx context.Context, providerId uint) error
	// Update Provider
	UpdateProvider(ctx context.Context, providerId uint, name string, displayName string, description string, homepageUrl string) error
	// Get Provider by hashIds
	GetProviderByHashIds(ctx context.Context, hashIds []string) ([]entity.Provider, error)
	// Get Provider By Id
	GetProviderById(ctx context.Context, providerId uint) (entity.Provider, error)
	// Get Provider By Name
	GetProviderByName(ctx context.Context, name string) (entity.Provider, error)
	// Update Hashid Provider
	UpdateHashIdProvider(ctx context.Context, providerId uint, hashId string) error
}
