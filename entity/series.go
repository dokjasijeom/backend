package entity

import (
	"gorm.io/gorm"
	"time"
)

type Series struct {
	Id          uint         `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement"`
	HashId      string       `gorm:"column:hash_id;type:varchar(255);null"`
	Title       string       `gorm:"column:title;type:varchar(255);not null"`
	Description string       `gorm:"column:description;type:text;null"`
	Thumbnail   string       `gorm:"column:thumbnail;type:varchar(255);not null"`
	ISBN        string       `gorm:"column:isbn;type:varchar(255);null"`
	ECNNumber   string       `gorm:"column:ecn_number;type:varchar(255);null"`
	SeriesType  string       `gorm:"column:series_type;type:varchar(255);not null"`
	PublishDays []PublishDay `gorm:"many2many:series_publish_days;"` // many to many relationship
	Episodes    []Episode    `gorm:"foreignKey:SeriesId"`
	Persons     []Person     `gorm:"many2many:series_persons;"` // many to many relationship
	Genres      []Genre      `gorm:"many2many:series_genres;"`  // many to many relationship
	PublisherId uint         `gorm:"column:publisher_id;type:int(11);null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
