package model

import (
	"github.com/dokjasijeom/backend/entity"
	"os"
)

type SeriesModel struct {
	Title       string            `form:"title"`
	Description string            `form:"description,omitempty"`
	Thumbnail   string            `form:"thumbnail,omitempty"`
	ISBN        string            `form:"isbn,omitempty"`
	ECN         string            `form:"ecn,omitempty"`
	Image       os.File           `form:"image,omitempty"`
	SeriesType  entity.SeriesType `form:"series_type"`
}
