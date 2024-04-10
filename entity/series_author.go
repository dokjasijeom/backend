package entity

import (
	"gorm.io/gorm"
	"time"
)

type SeriesAuthor struct {
	Id         uint   `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"-"`
	SeriesId   uint   `gorm:"column:series_id;type:int(11);not null;index;"`
	Series     Series `gorm:"foreignKey:SeriesId;references:Id" json:"series,omitempty"`
	PersonId   uint   `gorm:"column:person_id;type:int(11);not null;index;"`
	Person     Person `gorm:"foreignKey:PersonId;references:Id" json:"person,omitempty"`
	PersonType string `gorm:"column:person_type;type:varchar(255);null;index;" json:"personType,omitempty"`
	// person_type: author, illustrator, original_author
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (SeriesAuthor) TableName() string {
	return "series_authors"
}
