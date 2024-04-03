package entity

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	Id          uint   `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id,omitempty"`
	HashId      string `gorm:"column:hash_id;type:varchar(255);not null" json:"hashId,omitempty"`
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name,omitempty"`
	Description string `gorm:"column:description;type:text;null" json:"description,omitempty"`
	PersonType  string `gorm:"-:migration;->" json:"personType,omitempty"`

	AuthorSeries []Series `gorm:"many2many:series_authors;" json:"series,omitempty"` // many to many relationship

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Person) TableName() string {
	return "person"
}

func (person *Person) AfterDelete(tx *gorm.DB) (err error) {
	return tx.Model(person).Association("AuthorSeries").Clear()
}
