package model

import "github.com/dokjasijeom/backend/entity"

type UserRecordSeriesUpdateModel struct {
	ReadCompleted bool              `json:"readCompleted" default:"false"`
	Title         string            `json:"title,omitempty" validate:"required_with_all=Author Genre SeriesType"`
	Author        string            `json:"author,omitempty" validate:"required_with_all=Title Genre SeriesType"`
	Genre         string            `json:"genre,omitempty" validate:"required_with_all=Title Author SeriesType"`
	SeriesType    entity.SeriesType `json:"seriesType,omitempty" validate:"required_with_all=Title Author Genre, oneof=webnovel webtoon"`
}
