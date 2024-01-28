package model

type GenreRequestModel struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name"`
	ParentGenreId uint   `json:"parentGenreId"`
}
