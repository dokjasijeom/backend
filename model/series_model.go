package model

import (
	"github.com/dokjasijeom/backend/entity"
	"os"
)

type SeriesModel struct {
	Title             string              `form:"title"`
	Description       string              `form:"description,omitempty"`
	Thumbnail         string              `form:"thumbnail,omitempty"`
	ISBN              string              `form:"isbn,omitempty"`
	ECN               string              `form:"ecn,omitempty"`
	Image             os.File             `form:"image,omitempty"`
	SeriesType        entity.SeriesType   `form:"seriesType"`
	GenreId           uint                `form:"genreId,omitempty"`
	PersonId          uint                `form:"personId,omitempty"`
	AuthorId          uint                `form:"authorId,omitempty"`
	IllustratorId     uint                `form:"illustratorId,omitempty"`
	OriginalAuthorId  uint                `form:"originalAuthorId,omitempty"`
	PublisherId       uint                `form:"publisherId,omitempty"`
	ProviderId        uint                `form:"providerId,omitempty"`
	PublishDayId      uint                `form:"publishDayId,omitempty"`
	GenreIds          []uint              `form:"genreIds,omitempty"`
	PersonIds         []uint              `form:"personIds,omitempty"`
	PublisherIds      []uint              `form:"publisherIds,omitempty"`
	ProviderIds       []uint              `form:"providerIds,omitempty"`
	Providers         []ProviderLinkModel `form:"providers,omitempty"`
	PublishDayIds     []uint              `form:"publishDayIds,omitempty"`
	AuthorIds         []uint              `form:"authorIds,omitempty"`
	IllustratorIds    []uint              `form:"illustratorIds,omitempty"`
	OriginalAuthorIds []uint              `form:"originalAuthorIds,omitempty"`
	IsComplete        bool                `form:"isComplete,omitempty"`
}
