package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type UserService interface {
	// User Authentication
	AuthenticateUser(ctx context.Context, email, password string) (entity.User, error)
}
