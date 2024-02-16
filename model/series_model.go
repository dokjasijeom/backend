package model

import (
	"github.com/dokjasijeom/backend/entity"
	"os"
)

type SeriesModel struct {
	Title        string            `form:"title"`
	Description  string            `form:"description,omitempty"`
	Thumbnail    string            `form:"thumbnail,omitempty"`
	ISBN         string            `form:"isbn,omitempty"`
	ECN          string            `form:"ecn,omitempty"`
	Image        os.File           `form:"image,omitempty"`
	SeriesType   entity.SeriesType `form:"series_type"`
	GenreId      uint              `form:"genre_id,omitempty"`
	PersonId     uint              `form:"person_id,omitempty"`
	PublisherId  uint              `form:"publisher_id,omitempty"`
	ProviderId   uint              `form:"provider_id,omitempty"`
	PublishDayId uint              `form:"publish_day_id,omitempty"`
}
