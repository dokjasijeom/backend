package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) AuthenticateUser(ctx context.Context, email, password string) (entity.User, error) {
	userResult, err := userService.UserRepository.Authenticate(ctx, email)
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(password))
	if err != nil {
		panic(err)
	}

	return userResult, nil
}
