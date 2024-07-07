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

func NewGenreServiceImpl(genreRepository *repository.GenreRepository) service.GenreService {
	return &genreServiceImpl{GenreRepository: *genreRepository}
}

type genreServiceImpl struct {
	repository.GenreRepository
}

func (genreService *genreServiceImpl) CreateGenre(ctx context.Context, name string, genreType entity.GenreType, parentGenreId uint) (entity.Genre, error) {
	config := configuration.New()

	result, err := genreService.GenreRepository.CreateGenre(ctx, name, genreType, parentGenreId)
	hd := hashids.NewData()
	hd.Salt = config.Get("HASH_SALT_GENRE")
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(result.Id)})

	if e != "" {
		_ = genreService.GenreRepository.UpdateGenreHashId(ctx, result.Id, e)
	}
	if err != nil {
		exception.PanicLogging(err)
	}
	if result.Id == 0 {
		exception.PanicLogging("genre is empty")
	}
	return result, nil
}

func (genreService *genreServiceImpl) GetAllMainGenre(ctx context.Context) ([]entity.Genre, error) {
	result, err := genreService.GenreRepository.GetAllMainGenre(ctx)
	if err != nil {
		exception.PanicLogging(err)
		return nil, err
	}
	return result, nil
}

func (genreService *genreServiceImpl) GetAllSubGenre(ctx context.Context, parentGenreId uint) ([]entity.Genre, error) {
	result, err := genreService.GenreRepository.GetAllSubGenre(ctx, parentGenreId)
	if err != nil {
		exception.PanicLogging(err)
		return nil, err
	}
	return result, nil
}

func (genreService *genreServiceImpl) DeleteGenre(ctx context.Context, genreId uint) error {
	err := genreService.GenreRepository.DeleteGenre(ctx, genreId)
	if err != nil {
		exception.PanicLogging(err)
		return err
	}
	return nil
}

func (genreService *genreServiceImpl) UpdateGenre(ctx context.Context, genreId uint, name string) error {
	err := genreService.GenreRepository.UpdateGenre(ctx, genreId, name)
	if err != nil {
		exception.PanicLogging(err)
		return err
	}
	return nil
}

func (genreService *genreServiceImpl) GetGenreById(ctx context.Context, genreId uint) (entity.Genre, error) {
	result, err := genreService.GenreRepository.GetGenreById(ctx, genreId)
	if err != nil {
		exception.PanicLogging(err)
		return entity.Genre{}, err
	}
	return result, nil
}

func (genreService *genreServiceImpl) GetGenreByName(ctx context.Context, name string) (entity.Genre, error) {
	result, err := genreService.GenreRepository.GetGenreByName(ctx, name)
	if err != nil {
		exception.PanicLogging(err)
		return entity.Genre{}, err
	}
	return result, nil
}
