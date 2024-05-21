package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserRecordSeriesEpisode struct {
	Id                 uint             `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"-"`
	UserRecordSeries   UserRecordSeries `gorm:"foreignKey:UserRecordSeriesId;references:Id" json:"userRecordSeries,omitempty"`
	UserRecordSeriesId uint             `gorm:"column:user_record_series_id;type:int(11);not null" json:"userRecordSeriesId,omitempty"`
	EpisodeId          uint             `gorm:"column:episode_id;type:int(11);null" json:"episodeId,omitempty"`
	Episode            Episode          `gorm:"foreignKey:EpisodeId;references:Id" json:"episode,omitempty"`
	EpisodeNumber      uint             `gorm:"column:episode_number;type:int(11);null" json:"episodeNumber,omitempty"`
	Watched            bool             `gorm:"column:watched;type:tinyint(1);not null;default:0" json:"watched"`
	ProviderId         uint             `gorm:"column:provider_id;type:int(11);null" json:"providerId,omitempty"`
	Provider           Provider         `gorm:"foreignKey:ProviderId;references:Id" json:"provider,omitempty"`
	ProviderName       string           `gorm:"column:provider_name;type:varchar(255);null" json:"providerName,omitempty"`
	CreatedAt          time.Time        `json:"-"`
	DeletedAt          gorm.DeletedAt   `json:"-"`
}
