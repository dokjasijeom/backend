package model

import (
	"github.com/dokjasijeom/backend/entity"
	"os"
)

type SeriesModel struct {
	Title         string            `form:"title"`
	Description   string            `form:"description,omitempty"`
	Thumbnail     string            `form:"thumbnail,omitempty"`
	ISBN          string            `form:"isbn,omitempty"`
	ECN           string            `form:"ecn,omitempty"`
	Image         os.File           `form:"image,omitempty"`
	SeriesType    entity.SeriesType `form:"seriesType"`
	GenreId       uint              `form:"genreId,omitempty"`
	PersonId      uint              `form:"personId,omitempty"`
	PublisherId   uint              `form:"publisherId,omitempty"`
	ProviderId    uint              `form:"providerId,omitempty"`
	PublishDayId  uint              `form:"publishDayId,omitempty"`
	GenreIds      []uint            `form:"genreIds,omitempty"`
	PersonIds     []uint            `form:"personIds,omitempty"`
	PublisherIds  []uint            `form:"publisherIds,omitempty"`
	ProviderIds   []uint            `form:"providerIds,omitempty"`
	PublishDayIds []uint            `form:"publishDayIds,omitempty"`
	IsComplete    bool              `form:"isCompleted,omitempty"`
}
