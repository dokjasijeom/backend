package model

import "github.com/dokjasijeom/backend/entity"

type GenreRequestModel struct {
	Id            uint             `json:"id,omitempty"`
	Name          string           `json:"name"`
	GenreType     entity.GenreType `json:"genreType"`
	ParentGenreId uint             `json:"parentGenreId"`
}
