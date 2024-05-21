package model

type UserRecordSeriesEmptyModel struct {
	Id           uint   `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Author       string `json:"author,omitempty"`
	Genre        string `json:"genre,omitempty"`
	TotalEpisode uint   `json:"totalEpisode,omitempty"`
}
