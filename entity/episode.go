package entity

import (
	"gorm.io/gorm"
	"time"
)

type Episode struct {
	Id            uint     `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement"`
	Title         string   `gorm:"column:title;type:varchar(255);null"`
	EpisodeNumber uint     `gorm:"column:episode_number;type:int(11);not null"`
	Thumbnail     string   `gorm:"column:thumbnail;type:varchar(255);null"`
	Series        []Series `gorm:"many2many:series_episodes;"` // many to many relationship
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
