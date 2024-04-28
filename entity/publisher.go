package entity

import (
	"gorm.io/gorm"
	"time"
)

type Publisher struct {
	Id          uint           `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id,omitempty"`
	HashId      string         `gorm:"column:hash_id;type:varchar(255);not null" json:"hashId"`
	Name        string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description string         `gorm:"column:description;type:text;null" json:"description,omitempty"`
	HomepageUrl string         `gorm:"column:homepage_url;type:varchar(255);null" json:"homepageUrl,omitempty"`
	Series      []Series       `gorm:"many2many:series_publishers;" json:"series,omitempty"` // many to many relationship
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (publisher *Publisher) AfterDelete(tx *gorm.DB) (err error) {
	return tx.Model(publisher).Association("Series").Clear()
}
