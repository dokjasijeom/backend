package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewGenreRepositoryImpl(DB *gorm.DB) repository.GenreRepository {
	return &genreRepositoryImpl{DB: DB}
}

type genreRepositoryImpl struct {
	*gorm.DB
}

func (genreRepository *genreRepositoryImpl) CreateGenre(ctx context.Context, name string, genreType entity.GenreType, parentGenreId uint) (entity.Genre, error) {
	var genreResult entity.Genre
	genreResult.Name = name
	genreResult.GenreType = genreType

	if parentGenreId != 0 {
		genreResult.ParentGenreId = parentGenreId
	}

	result := genreRepository.DB.WithContext(ctx).Create(&genreResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		genreResult.Id = 0
		return genreResult, result.Error
	}
	return genreResult, nil
}

func (genreRepository *genreRepositoryImpl) GetAllMainGenre(ctx context.Context) ([]entity.Genre, error) {
	var genreResult []entity.Genre
	result := genreRepository.DB.WithContext(ctx).Where("parent_genre_id = 0").Find(&genreResult)

	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return nil, result.Error
	}

	return genreResult, nil
}

func (genreRepository *genreRepositoryImpl) GetAllSubGenre(ctx context.Context, parentGenreId uint) ([]entity.Genre, error) {
	var subGenreResult []entity.Genre
	result := genreRepository.DB.WithContext(ctx).Where("parent_genre_id = ?", parentGenreId).Find(&subGenreResult)

	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return nil, result.Error
	}

	return subGenreResult, nil
}

func (genreRepository *genreRepositoryImpl) DeleteGenre(ctx context.Context, genreId uint) error {
	result := genreRepository.DB.WithContext(ctx).Delete(&entity.Genre{}, genreId)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}

	return nil
}

func (genreRepository *genreRepositoryImpl) UpdateGenre(ctx context.Context, genreId uint, name string) error {
	result := genreRepository.DB.WithContext(ctx).Model(&entity.Genre{}).Where("id = ?", genreId).Update("name", name)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}

	return nil
}

func (genreRepository *genreRepositoryImpl) GetGenreById(ctx context.Context, genreId uint) (entity.Genre, error) {
	var genreResult entity.Genre
	result := genreRepository.DB.WithContext(ctx).Where("id = ?", genreId).Find(&genreResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return entity.Genre{}, result.Error
	}

	return genreResult, nil
}

func (genreRepository *genreRepositoryImpl) GetGenreByName(ctx context.Context, name string) (entity.Genre, error) {
	var genreResult entity.Genre
	result := genreRepository.DB.WithContext(ctx).Where("name = ?", name).Find(&genreResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return entity.Genre{}, result.Error
	}

	return genreResult, nil
}

func (genreRepository *genreRepositoryImpl) UpdateGenreHashId(ctx context.Context, genreId uint, hashId string) error {
	result := genreRepository.DB.WithContext(ctx).Model(&entity.Genre{}).Where("id = ?", genreId).Update("hash_id", hashId)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return result.Error
	}

	return nil
}
