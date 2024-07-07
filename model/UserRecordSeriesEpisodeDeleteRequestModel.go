package model

type UserRecordSeriesEpisodeDeleteRequestModel struct {
	UserRecordSeriesId uint   `json:"userRecordSeriesId,required"`
	RecordIds          []uint `json:"recordIds,required"`
}
