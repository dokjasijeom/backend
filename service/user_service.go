package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type UserService interface {
	// User Authentication
	AuthenticateUser(ctx context.Context, email, password string) (entity.User, error)
	// User Create
	CreateUser(ctx context.Context, email, password, comparePassword string) (entity.User, error)
	// Get User By Email
	GetUserByEmail(ctx context.Context, email string) entity.User
	// Get User By Email And Series
	GetUserByEmailAndSeries(ctx context.Context, email string) entity.User
}
