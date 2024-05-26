package model

type UserRecordSeriesEmptyModel struct {
	Id           uint   `json:"id,omitempty" validate:"required_without_all=Title Author Genre TotalEpisode"`
	Title        string `json:"title,omitempty" validate:"required_with_all=Author Genre TotalEpisode"`
	Author       string `json:"author,omitempty" validate:"required_with_all=Title Genre TotalEpisode"`
	Genre        string `json:"genre,omitempty" validate:"required_with_all=Title Author TotalEpisode"`
	TotalEpisode uint   `json:"totalEpisode,omitempty" validate:"required_with_all=Title Author Genre, min=1"`
}
