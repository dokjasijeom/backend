package entity

import (
	"gorm.io/gorm"
	"time"
)

type SeriesProvider struct {
	Id         uint     `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement"`
	SeriesId   uint     `gorm:"column:series_id;type:int(11);not null;index"`
	Series     Series   `gorm:"foreignKey:SeriesId;references:Id" json:"series,omitempty"`
	ProviderId uint     `gorm:"column:provider_id;type:int(11);not null;index"`
	Provider   Provider `gorm:"foreignKey:ProviderId;references:Id" json:"provider,omitempty"`
	Link       string   `gorm:"column:link;type:varchar(255);null" json:"link,omitempty"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
