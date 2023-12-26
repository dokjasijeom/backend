package impl

import (
	"context"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"github.com/speps/go-hashids/v2"
)

func NewUserServiceImpl(userRepository *repository.UserRepository, userRoleRepository *repository.UserRoleRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository, UserRoleRepository: *userRoleRepository}
}

type userServiceImpl struct {
	repository.UserRepository
	repository.UserRoleRepository
}

func (userService *userServiceImpl) CreateUser(ctx context.Context, email, password, comparePassword string) (entity.User, error) {
	if password != comparePassword {
		exception.PanicLogging("password and compare password is not same")
	}
	config := configuration.New()

	result, err := userService.UserRepository.CreateUser(email, password)
	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(result.Id)})

	if e != "" {
		_ = userService.UpdateUserHashId(ctx, email, e)
	}
	if err != nil {
		exception.PanicLogging(err)
	}
	if result.Id == 0 {
		exception.PanicLogging("user is empty")
	} else {
		_, err := userService.UserRoleRepository.CreateUserRole(ctx, result.Id, 2)
		if err != nil {
			exception.PanicLogging(err)
		}
	}
	return result, nil
}

func (userService *userServiceImpl) AuthenticateUser(ctx context.Context, email, password string) (entity.User, error) {
	userResult, err := userService.UserRepository.Authenticate(ctx, email)
	if err != nil {
		panic(err)
	}
	match, err := userService.UserRepository.CompareHashAndPassword(password, userResult.Password)
	if err != nil {
		panic(err)
	}

	if !match {
		exception.PanicLogging("password is not matched")
	} else {
		return userResult, nil
	}

	return userResult, nil
}

func (userService *userServiceImpl) GetUserByEmail(ctx context.Context, email string) entity.User {
	userResult, _ := userService.UserRepository.GetUserByEmail(ctx, email)
	return userResult
}
