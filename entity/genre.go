package entity

import (
	"gorm.io/gorm"
	"time"
)

type GenreType string

const (
	Common   GenreType = "common"
	Webtoon  GenreType = "webtoon"
	Webnovel GenreType = "webnovel"
)

type Genre struct {
	Id            uint      `gorm:"primaryKey;column:id;type:int(11);autoIncrement" json:"id,omitempty"`
	HashId        string    `gorm:"column:hash_id;type:varchar(255);not null;unique" json:"hashId"`
	Name          string    `gorm:"column:name;type:varchar(255);unique;not null" json:"name"`
	GenreType     GenreType `gorm:"column:genre_type;type:varchar(20);not null" json:"genreType,omitempty"`
	ParentGenreId uint      `gorm:"column:parent_genre_id;type:int(11);null" json:"parentGenreId,omitempty"`
	// ParentGenreId: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
