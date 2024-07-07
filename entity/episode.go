package entity

import (
	"gorm.io/gorm"
	"time"
)

type Episode struct {
	Id            uint           `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id"` //  form:"id" query:"id" validate:"required"
	Title         string         `gorm:"column:title;type:varchar(255);null" json:"title,omitempty"`         // form:"title" query:"title" validate:"required"
	EpisodeNumber uint           `gorm:"column:episode_number;type:int(11);not null" json:"episodeNumber"`   // form:"episodeNumber" query:"episodeNumber" validate:"required"
	Thumbnail     string         `gorm:"column:thumbnail;type:varchar(255);null" json:"thumbnail,omitempty"` // form:"thumbnail" query:"thumbnail" validate:"required"
	Series        []Series       `gorm:"many2many:series_episodes;" json:"series,omitempty"`                 // many to many relationship
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
