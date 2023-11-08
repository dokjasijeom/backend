package impl

import (
	"context"
	"errors"
	"github.com/dokjasijeom/backend/entity"
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

func (userRepository *userRepositoryImpl) GetUserByEmail(email string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) GetUserByEmailAndPassword(email, password string) error {
	//TODO implement me
	panic("implement me")
}

func (userRepository *userRepositoryImpl) CreateUser(email, password string) error {
	//TODO implement me
	panic("implement me")
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
