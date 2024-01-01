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
