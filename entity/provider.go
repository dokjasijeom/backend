package entity

import (
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	Id          uint     `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement"`
	HashId      string   `gorm:"column:hash_id;type:varchar(255);not null"`
	Name        string   `gorm:"column:name;type:varchar(255);not null"`
	DisplayName string   `gorm:"column:display_name;type:varchar(255);not null"`
	Description string   `gorm:"column:description;type:text;null"`
	HomepageUrl string   `gorm:"column:homepage_url;type:varchar(255);null"`
	Series      []Series `gorm:"many2many:series_providers;"` // many to many relationship
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
