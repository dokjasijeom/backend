package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewUserRoleRepositoryImpl(DB *gorm.DB) repository.UserRoleRepository {
	return &userRoleRepositoryImpl{DB: DB}
}

type userRoleRepositoryImpl struct {
	*gorm.DB
}

func (userRoleRepository *userRoleRepositoryImpl) CreateUserRole(ctx context.Context, userId uint, roleId uint) (entity.UserRole, error) {
	var userRoleResult entity.UserRole
	userRoleResult.UserId = userId
	userRoleResult.RoleId = roleId
	result := userRoleRepository.DB.WithContext(ctx).Create(&userRoleResult)
	if result.RowsAffected == 0 {
		return entity.UserRole{}, nil
	}
	return userRoleResult, nil
}

func (userRoleRepository *userRoleRepositoryImpl) GetUserRoleByUserId(userId uint) (entity.UserRole, error) {
	var userRoleResult entity.UserRole
	err := userRoleRepository.DB.Where("user_id = ?", userId).Find(&userRoleResult)
	if err != nil {
		exception.PanicLogging(err)
	}
	return userRoleResult, nil
}

func (userRoleRepository *userRoleRepositoryImpl) GetUserRoleByUserIdAndRole(userId uint, roleId uint) (entity.UserRole, error) {
	//TODO implement me
	panic("implement me")
}

func (userRoleRepository *userRoleRepositoryImpl) UpdateUserRoleByUserId(userId uint) error {
	//TODO implement me
	panic("implement me")
}

func (userRoleRepository *userRoleRepositoryImpl) DeleteUserRoleByUserId(userId uint) error {
	//TODO implement me
	panic("implement me")
}
