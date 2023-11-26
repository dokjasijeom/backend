package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type RoleService interface {
	// Get role by name
	GetRoleByName(ctx context.Context, name string) (entity.Role, error)
	// Create new role
	CreateRole(ctx context.Context, name string) (entity.Role, error)
}
