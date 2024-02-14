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

func NewPersonServiceImpl(personRepository *repository.PersonRepository) service.PersonService {
	return &personService{PersonRepository: *personRepository}
}

type personService struct {
	repository.PersonRepository
}

func (personService *personService) CreatePerson(ctx context.Context, name string) (entity.Person, error) {
	config := configuration.New()

	result, err := personService.PersonRepository.CreatePerson(ctx, name)
	if err != nil {
		exception.PanicLogging(err)
	}
	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT_PERSON")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(result.Id)})

	if e != "" {
		_ = personService.UpdatePersonHashId(ctx, result, e)
		result.HashId = e
	}

	return result, nil
}

func (personService *personService) UpdatePerson(ctx context.Context, id uint, name string) (entity.Person, error) {
	result, err := personService.PersonRepository.UpdatePerson(ctx, id, name)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}

func (personService *personService) UpdatePersonHashId(ctx context.Context, person entity.Person, hashId string) error {
	err := personService.PersonRepository.UpdatePersonHashId(ctx, person, hashId)
	if err != nil {
		exception.PanicLogging(err)
	}
	return nil
}

func (personService *personService) GetPersonById(ctx context.Context, id uint) (entity.Person, error) {
	result, err := personService.PersonRepository.GetPersonById(ctx, id)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}

func (personService *personService) GetPersonByHashId(ctx context.Context, hashId string) (entity.Person, error) {
	result, err := personService.PersonRepository.GetPersonByHashId(ctx, hashId)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}

func (personService *personService) GetPersonByName(ctx context.Context, name string) (entity.Person, error) {
	result, err := personService.PersonRepository.GetPersonByName(ctx, name)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}

func (personService *personService) DeletePersonById(ctx context.Context, id uint) (entity.Person, error) {
	result, err := personService.PersonRepository.DeletePersonById(ctx, id)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}

func (personService *personService) DeletePersonByHashId(ctx context.Context, hashId string) (entity.Person, error) {
	result, err := personService.PersonRepository.DeletePersonByHashId(ctx, hashId)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}

func (personService *personService) GetAllPerson(ctx context.Context) ([]entity.Person, error) {
	result, err := personService.PersonRepository.GetAllPerson(ctx)
	if err != nil {
		exception.PanicLogging(err)
	}
	return result, nil
}
