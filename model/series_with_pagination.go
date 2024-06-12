package model

import "github.com/dokjasijeom/backend/entity"

type SeriesWithPagination struct {
	Series     []entity.Series `json:"series"`
	Pagination Pagination      `json:"pagination"`
}
