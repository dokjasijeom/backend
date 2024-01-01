package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PersonService interface {
	// create new person
	CreatePerson(ctx context.Context, name string) (entity.Person, error)
}
