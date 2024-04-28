package entity

import (
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	Id             uint             `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id,omitempty"`
	HashId         string           `gorm:"column:hash_id;type:varchar(255);not null" json:"hashId"`
	Name           string           `gorm:"column:name;type:varchar(255);not null" json:"name"`
	DisplayName    string           `gorm:"column:display_name;type:varchar(255);not null" json:"displayName"`
	Description    string           `gorm:"column:description;type:text;null" json:"description"`
	HomepageUrl    string           `gorm:"column:homepage_url;type:varchar(255);null" json:"homepageUrl"`
	ProviderSeries []SeriesProvider `json:"series,omitempty"` // many to many relationship
	Link           string           `json:"link,omitempty"`
	//Series      []Series       `gorm:"many2many:series_providers;" json:"series"` // many to many relationship
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (provider *Provider) AfterDelete(tx *gorm.DB) (err error) {
	return tx.Model(provider).Association("Series").Clear()
}
