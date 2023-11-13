package impl

import (
	"context"
	"errors"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) GetAllUsers() error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) GetUserByEmail(email string) entity.User {
	var userResult entity.User
	result := userRepository.DB.Where("user.email = ?", email).Find(&userResult)
	if result.RowsAffected == 0 {
		exception.PanicLogging("user not found")
	}

	return userResult
}

func (userRepository *userRepositoryImpl) GetUserByEmailAndPassword(email, password string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) CreateUser(email, password string) error {
	var userResult entity.User
	userResult.Email = email
	userResult.Password = password
	result := userRepository.DB.Create(&userResult)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userRepository *userRepositoryImpl) UpdateUserByEmail(email string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) DeleteUserByEmail(email string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) Authenticate(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("user.email = ?", email).Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
