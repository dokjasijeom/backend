package entity

import (
	"gorm.io/gorm"
	"time"
)

type SeriesDailyView struct {
	Id        uint      `gorm:"primaryKey;column:id;type:int(11);autoIncrement;not null" json:"id,omitempty"`
	SeriesId  uint      `gorm:"column:series_id;type:int(11);not null;index;" json:"seriesId,omitempty"`
	ViewCount uint      `gorm:"column:view_count;type:int(11);not null;default:0;index" json:"viewCount,omitempty"`
	ViewDate  time.Time `gorm:"column:view_date;type:date;not null;index" json:"viewDate,omitempty"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (sdv *SeriesDailyView) AfterCreate(tx *gorm.DB) (err error) {
	var series Series
	tx.Model(&Series{}).Where("id = ?", sdv.SeriesId).First(&series)
	series.ViewCount++
	tx.Updates(&series)

	return nil
}
