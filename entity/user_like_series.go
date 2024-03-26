package entity

import "gorm.io/gorm"

type UserLikeSeries struct {
	UserId   uint `gorm:"primaryKey;column:user_id;type:int(11);not null"`
	SeriesId uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
}

func (UserLikeSeries) TableName() string {
	return "user_like_series"
}

func (uls *UserLikeSeries) AfterCreate(tx *gorm.DB) (err error) {
	var ulsc UserLikeSeriesCount
	tx.Model(&ulsc).Where("series_id = ?", uls.SeriesId).First(&ulsc)
	if ulsc != (UserLikeSeriesCount{}) {
		ulsc.Count++
		tx.Save(&ulsc)
	} else {
		ulsc.SeriesId = uls.SeriesId
		ulsc.Count = 1
		tx.Create(&ulsc)
	}

	var series Series
	tx.Model(&series).Where("id = ?", uls.SeriesId).First(&series)
	series.LikeCount++
	tx.Save(&series)

	return nil
}

func (uls *UserLikeSeries) AfterDelete(tx *gorm.DB) (err error) {
	var ulsc UserLikeSeriesCount
	tx.Model(&ulsc).Where("series_id = ?", uls.SeriesId).First(&ulsc)
	if ulsc != (UserLikeSeriesCount{}) {
		ulsc.Count--
		tx.Save(&ulsc)
	}

	var series Series
	tx.Model(&series).Where("id = ?", uls.SeriesId).First(&series)
	series.LikeCount--
	tx.Save(&series)

	return nil
}
