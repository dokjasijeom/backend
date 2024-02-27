package entity

import (
	"gorm.io/gorm"
	"time"
)

type SeriesType string

const (
	WebNovel SeriesType = "webnovel"
	WebToon  SeriesType = "webtoon"
)

type Series struct {
	Id          uint         `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id,omitempty"`
	HashId      string       `gorm:"column:hash_id;type:varchar(255);null" json:"hashId"`
	Title       string       `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string       `gorm:"column:description;type:text;null" json:"description"`
	Thumbnail   string       `gorm:"column:thumbnail;type:varchar(255);not null" json:"thumbnail"`
	ISBN        string       `gorm:"column:isbn;type:varchar(255);null" json:"isbn"`
	ECN         string       `gorm:"column:ecn;type:varchar(255);null" json:"ecn"`
	SeriesType  SeriesType   `gorm:"column:series_type;type:varchar(20);not null" json:"seriesType"`
	PublishDays []PublishDay `gorm:"many2many:series_publish_days;" json:"publishDays"` // many to many relationship
	Episodes    []Episode    `gorm:"many2many:series_episodes;" json:"episodes"`        // many to many relationship
	//Persons     []Person     `gorm:"many2many:series_persons;" json:"people"`           // many to many relationship
	Genres      []Genre    `gorm:"many2many:series_genres;" json:"genres"`   // many to many relationship
	Authors     []Person   `gorm:"many2many:series_authors;" json:"authors"` // many to many relationship
	PublisherId uint       `gorm:"column:publisher_id;type:int(11);null" json:"publisherId"`
	Publisher   Publisher  `gorm:"foreignKey:PublisherId" json:"publisher"`
	Providers   []Provider `gorm:"many2many:series_providers;" json:"providers"` // many to many relationship

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
