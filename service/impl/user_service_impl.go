package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
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

func (userService *userServiceImpl) CreateUser(ctx context.Context, email, password, comparePassword string) error {
	if password != comparePassword {
		exception.PanicLogging("password and compare password is not same")
	}

	err := userService.UserRepository.CreateUser(email, password)
	if err != nil {
		exception.PanicLogging(err)
	}
	return nil
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

func (userService *userServiceImpl) GetUserByEmail(ctx context.Context, email string) entity.User {
	userResult, err := userService.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		panic(err)
	}
	return userResult
}
