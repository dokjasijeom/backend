package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewRoleServiceImpl(roleRepository *repository.RoleRepository) service.RoleService {
	return &roleServiceImpl{RoleRepository: *roleRepository}
}

type roleServiceImpl struct {
	repository.RoleRepository
}

func (roleService *roleServiceImpl) CreateRole(ctx context.Context, name string) (entity.Role, error) {
	result, err := roleService.RoleRepository.CreateRole(ctx, name)
	if err != nil {
		exception.PanicLogging(err)
		panic(err)
	}
	return result, nil
}

func (roleService *roleServiceImpl) GetRoleByName(ctx context.Context, name string) (entity.Role, error) {
	result, err := roleService.RoleRepository.GetRoleByName(ctx, name)
	if err != nil {
		panic(err)
	}
	return result, nil
}
