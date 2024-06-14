package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserRecordSeries struct {
	Id                 uint                      `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id"`
	UserId             uint                      `gorm:"column:user_id;type:int(11);not null,index" json:"-"`
	User               User                      `gorm:"foreignKey:UserId;references:Id" json:"-"`
	SeriesId           uint                      `gorm:"column:series_id;type:int(11);null,index" json:"-"`
	Series             *Series                   `gorm:"foreignKey:SeriesId;references:Id" json:"series,omitempty"`
	Title              string                    `gorm:"column:title;type:varchar(255);null" json:"title,omitempty"`
	Author             string                    `gorm:"column:author;type:varchar(255);null" json:"author,omitempty"`
	Genre              string                    `gorm:"column:genre;type:varchar(255);null" json:"genre,omitempty"`
	SeriesType         SeriesType                `gorm:"column:series_type;type:varchar(255);null" json:"seriesType"`
	RecordEpisodeCount uint                      `gorm:"column:record_episode_count;type:int(11);null;default:0" json:"recordEpisodeCount"`
	RecordEpisodes     []UserRecordSeriesEpisode `json:"recordEpisodes"`
	RecordProviders    []string                  `gorm:"-" json:"recordProviders"`
	CreatedAt          time.Time                 `json:"-"`
	DeletedAt          gorm.DeletedAt            `json:"-"`
}
