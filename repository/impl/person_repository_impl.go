package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewPersonRepositoryImpl(DB *gorm.DB) repository.PersonRepository {
	return &personRepositoryImpl{DB: DB}
}

type personRepositoryImpl struct {
	*gorm.DB
}

func (personRepository *personRepositoryImpl) GetPersonById(ctx context.Context, id uint) (entity.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (personRepository *personRepositoryImpl) GetPersonByHashId(ctx context.Context, hashId string) (entity.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (personRepository *personRepositoryImpl) GetPersonByName(ctx context.Context, name string) (entity.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (personRepository *personRepositoryImpl) CreatePerson(ctx context.Context, name string) (entity.Person, error) {
	var personResult entity.Person
	personResult.Name = name
	result := personRepository.DB.WithContext(ctx).Create(&personResult)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
}

func (personRepository *personRepositoryImpl) UpdatePerson(ctx context.Context, person entity.Person) (entity.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (personRepository *personRepositoryImpl) UpdatePersonHashId(ctx context.Context, person entity.Person, hashId string) error {
	//TODO implement me
	panic("implement me")
}

func (personRepository *personRepositoryImpl) DeletePersonById(ctx context.Context, id uint) (entity.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (personRepository *personRepositoryImpl) DeletePersonByHashId(ctx context.Context, hashId string) (entity.Person, error) {
	//TODO implement me
	panic("implement me")
}
