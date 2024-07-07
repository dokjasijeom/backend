package model

type UserRecordSeriesEmptyModel struct {
	Id         uint   `json:"id,omitempty" validate:"required_without_all=Title Author Genre SeriesType"`
	Title      string `json:"title,omitempty" validate:"required_with_all=Author Genre SeriesType"`
	Author     string `json:"author,omitempty" validate:"required_with_all=Title Genre SeriesType"`
	Genre      string `json:"genre,omitempty" validate:"required_with_all=Title Author SeriesType"`
	SeriesType string `json:"seriesType,omitempty" validate:"required_with_all=Title Author Genre, oneof=webnovel webtoon"`
}
