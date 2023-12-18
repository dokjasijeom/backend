package entity

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	Id          uint   `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement"`
	HashId      string `gorm:"column:hash_id;type:varchar(255);not null"`
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text;not null"`

	Series []Series `gorm:"many2many:series_persons;"` // many to many relationship

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
