package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type UserRoleRepository interface {
	// Get user role by user id
	GetUserRoleByUserId(userId uint) (entity.UserRole, error)
	// Get user role by user id and role
	GetUserRoleByUserIdAndRole(userId uint, roleId uint) (entity.UserRole, error)
	// Create new user role
	CreateUserRole(ctx context.Context, userId uint, roleId uint) (entity.UserRole, error)
	// Update user role by user id
	UpdateUserRoleByUserId(userId uint) error
	// Delete user role by user id
	DeleteUserRoleByUserId(userId uint) error
}
