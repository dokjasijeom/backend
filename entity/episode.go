package entity

import (
	"gorm.io/gorm"
	"time"
)

type Episode struct {
	Id            uint   `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement"`
	HashId        string `gorm:"column:hash_id;type:varchar(255);not null"`
	Title         string `gorm:"column:title;type:varchar(255);not null"`
	SeriesId      uint   `gorm:"column:series_id;type:int(11);not null"`
	EpisodeNumber uint   `gorm:"column:episode_number;type:int(11);not null"`
	Thumbnail     string `gorm:"column:thumbnail;type:varchar(255);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
