package service

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
)

type GenreService interface {
	// Create New Genre
	CreateGenre(ctx context.Context, name string, genreType entity.GenreType, parentGenreId uint) (entity.Genre, error)
	// Get All Main Genre
	GetAllMainGenre(ctx context.Context) ([]entity.Genre, error)
	// Get All Sub Genre
	GetAllSubGenre(ctx context.Context, parentGenreId uint) ([]entity.Genre, error)
	// Delete Genre
	DeleteGenre(ctx context.Context, genreId uint) error
	// Update Genre
	UpdateGenre(ctx context.Context, genreId uint, name string) error
	// Get Genre By Id
	GetGenreById(ctx context.Context, genreId uint) (entity.Genre, error)
	// Get Genre By Name
	GetGenreByName(ctx context.Context, name string) (entity.Genre, error)
}
