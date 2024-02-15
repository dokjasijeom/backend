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
	Id            uint      `gorm:"primaryKey;column:id;type:int(11);autoIncrement"`
	HashId        string    `gorm:"column:hash_id;type:varchar(255);not null;unique"`
	Name          string    `gorm:"column:name;type:varchar(255);unique;not null"`
	GenreType     GenreType `gorm:"column:genre_type;type:enum('common', 'webtoon', 'webnovel');not null"`
	ParentGenreId uint      `gorm:"column:parent_genre_id;type:int(11);null"`
	// ParentGenreId: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
