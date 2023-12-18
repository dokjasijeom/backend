package entity

import (
	"gorm.io/gorm"
	"time"
)

type Genre struct {
	Id          uint   `gorm:"primaryKey;column:id;type:int(11)"`
	Name        string `gorm:"column:name;type:varchar(255);unique;not null"`
	GenreNumber int    `gorm:"column:genre_number;type:int(11);unique;null"`
	// GenreNumber: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	ParentGenereId uint `gorm:"column:parent_genre_id;type:int(11);null"`
	// ParentGenreId: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
