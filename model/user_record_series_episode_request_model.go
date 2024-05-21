package model

type UserRecordSeriesEpisodeRequestModel struct {
	UserRecordSeriesId uint   `json:"userRecordSeriesId,required"`
	ProviderName       string `json:"providerName,required"`
	From               uint   `json:"from,omitempty"`
	To                 uint   `json:"to,required"`
}
