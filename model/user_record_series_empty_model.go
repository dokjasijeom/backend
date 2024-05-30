package model

type UserRecordSeriesEmptyModel struct {
	Id           uint   `json:"id,omitempty" validate:"required_without_all=Title Author Genre TotalEpisode SeriesType"`
	Title        string `json:"title,omitempty" validate:"required_with_all=Author Genre TotalEpisode SeriesType"`
	Author       string `json:"author,omitempty" validate:"required_with_all=Title Genre TotalEpisode SeriesType"`
	Genre        string `json:"genre,omitempty" validate:"required_with_all=Title Author TotalEpisode SeriesType"`
	TotalEpisode uint   `json:"totalEpisode,omitempty" validate:"required_with_all=Title Author Genre SeriesType, min=1"`
	SeriesType   string `json:"seriesType,omitempty" validate:"required_with_all=Title Author Genre TotalEpisode, oneof=webnovel webtoon"`
}
