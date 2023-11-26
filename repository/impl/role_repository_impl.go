package impl

import (
	"context"
	"errors"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewRoleRepositoryImpl(DB *gorm.DB) repository.RoleRepository {
	return &roleRepositoryImpl{DB: DB}
}

type roleRepositoryImpl struct {
	*gorm.DB
}

func (roleRepository *roleRepositoryImpl) GetRoleByName(ctx context.Context, name string) (entity.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (roleRepository *roleRepositoryImpl) CreateRole(ctx context.Context, name string) (entity.Role, error) {
	var roleResult entity.Role
	roleResult.Role = name
	result := roleRepository.DB.WithContext(ctx).Create(&roleResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging(errors.New("already exist"))
		return entity.Role{}, errors.New("already exist")
	}
	return roleResult, nil
}
