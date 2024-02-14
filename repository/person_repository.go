package repository

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type PersonRepository interface {
	// get person by id
	GetPersonById(ctx context.Context, id uint) (entity.Person, error)
	// get person by hash id
	GetPersonByHashId(ctx context.Context, hashId string) (entity.Person, error)
	// get person by name
	GetPersonByName(ctx context.Context, name string) (entity.Person, error)
	// create new person
	CreatePerson(ctx context.Context, name string) (entity.Person, error)
	// update person
	UpdatePerson(ctx context.Context, id uint, name string) (entity.Person, error)
	// update person's hashId
	UpdatePersonHashId(ctx context.Context, person entity.Person, hashId string) error
	// delete person by id
	DeletePersonById(ctx context.Context, id uint) (entity.Person, error)
	// delete person by hash id
	DeletePersonByHashId(ctx context.Context, hashId string) (entity.Person, error)
	// get all person
	GetAllPerson(ctx context.Context) ([]entity.Person, error)
}
