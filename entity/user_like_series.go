package entity

import (
	"gorm.io/gorm"
)

type UserLikeSeries struct {
	UserId   uint `gorm:"primaryKey;column:user_id;type:int(11);not null"`
	SeriesId uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
}

func (uls *UserLikeSeries) AfterCreate(tx *gorm.DB) (err error) {
	var ulsc UserLikeSeriesCount
	tx.Model(&UserLikeSeriesCount{}).Where("series_id = ?", uls.SeriesId).First(&ulsc)
	if ulsc != (UserLikeSeriesCount{}) {
		ulsc.Count++
		tx.Updates(&ulsc)
	} else {
		ulsc.SeriesId = uls.SeriesId
		ulsc.Count = 1
		tx.Create(&ulsc)
	}

	var series Series
	tx.Model(&Series{}).Where("id = ?", uls.SeriesId).First(&series)
	series.LikeCount++
	tx.Updates(&series)

	return nil
}

func (uls *UserLikeSeries) BeforeDelete(tx *gorm.DB) (err error) {
	var ulsc UserLikeSeriesCount
	tx.Model(&UserLikeSeriesCount{}).Where("series_id = ?", uls.SeriesId).First(&ulsc)
	if ulsc != (UserLikeSeriesCount{}) {
		ulsc.Count--
		tx.Updates(&ulsc)
	}

	var series Series
	tx.Model(&Series{}).Where("id = ?", uls.SeriesId).First(&series)
	series.LikeCount = series.LikeCount - 1
	tx.Updates(&series)

	return nil
}
