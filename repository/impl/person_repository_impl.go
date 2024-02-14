package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
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
	var personResult entity.Person
	result := personRepository.DB.WithContext(ctx).First(&personResult, id)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
}

func (personRepository *personRepositoryImpl) GetPersonByHashId(ctx context.Context, hashId string) (entity.Person, error) {
	var personResult entity.Person
	result := personRepository.DB.WithContext(ctx).Where("hash_id = ?", hashId).First(&personResult)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
}

func (personRepository *personRepositoryImpl) GetPersonByName(ctx context.Context, name string) (entity.Person, error) {
	var personResult entity.Person
	result := personRepository.DB.WithContext(ctx).Where("name = ?", name).First(&personResult)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
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

func (personRepository *personRepositoryImpl) UpdatePerson(ctx context.Context, id uint, name string) (entity.Person, error) {
	var personResult entity.Person
	result := personRepository.DB.WithContext(ctx).First(&personResult, id)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	personResult.Name = name
	result = personRepository.DB.WithContext(ctx).Save(&personResult)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
}

func (personRepository *personRepositoryImpl) UpdatePersonHashId(ctx context.Context, person entity.Person, hashId string) error {
	result := personRepository.DB.WithContext(ctx).Model(&entity.Person{}).Where("id = ?", person.Id).Updates(map[string]interface{}{"hash_id": hashId})

	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}

	return nil
}

func (personRepository *personRepositoryImpl) DeletePersonById(ctx context.Context, id uint) (entity.Person, error) {
	var personResult entity.Person
	result := personRepository.DB.WithContext(ctx).First(&personResult, id)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	result = personRepository.DB.WithContext(ctx).Delete(&personResult, id)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
}

func (personRepository *personRepositoryImpl) DeletePersonByHashId(ctx context.Context, hashId string) (entity.Person, error) {
	var personResult entity.Person
	result := personRepository.DB.WithContext(ctx).Where("hash_id = ?", hashId).First(&personResult)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	result = personRepository.DB.WithContext(ctx).Delete(&personResult, hashId)
	if result.RowsAffected == 0 {
		return entity.Person{}, nil
	}
	return personResult, nil
}

func (personRepository *personRepositoryImpl) GetAllPerson(ctx context.Context) ([]entity.Person, error) {
	var personResult []entity.Person
	result := personRepository.DB.WithContext(ctx).Find(&personResult)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return personResult, nil
}
