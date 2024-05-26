package entity

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type UserRecordSeriesEpisode struct {
	Id                 uint             `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id"`
	UserRecordSeries   UserRecordSeries `gorm:"foreignKey:UserRecordSeriesId;references:Id" json:"-"`
	UserRecordSeriesId uint             `gorm:"column:user_record_series_id;type:int(11);not null" json:"-"`
	EpisodeId          uint             `gorm:"column:episode_id;type:int(11);null" json:"-"`
	Episode            *Episode         `gorm:"foreignKey:EpisodeId;references:Id" json:"-"`
	EpisodeNumber      uint             `gorm:"column:episode_number;type:int(11);null" json:"episodeNumber,omitempty"`
	Watched            bool             `gorm:"column:watched;type:tinyint(1);not null;default:0" json:"watched"`
	ProviderId         uint             `gorm:"column:provider_id;type:int(11);null" json:"-"`
	Provider           *Provider        `gorm:"foreignKey:ProviderId;references:Id" json:"-"`
	ProviderName       string           `gorm:"column:provider_name;type:varchar(255);null" json:"providerName,omitempty"`
	CreatedAt          time.Time        `json:"-"`
	DeletedAt          gorm.DeletedAt   `json:"-"`
}

// AfterCreate
func (userRecordSeriesEpisode *UserRecordSeriesEpisode) AfterCreate(tx *gorm.DB) (err error) {
	var urs UserRecordSeries
	tx.Model(&UserRecordSeries{}).First(&urs, userRecordSeriesEpisode.UserRecordSeriesId)
	urs.RecordEpisodeCount = urs.RecordEpisodeCount + 1
	tx.Updates(&urs)

	return nil
}

// BeforeDelete
func (userRecordSeriesEpisode *UserRecordSeriesEpisode) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println(userRecordSeriesEpisode)
	var urs UserRecordSeries
	tx.Model(&UserRecordSeries{}).Where("id = ?", userRecordSeriesEpisode.UserRecordSeriesId).First(&urs)
	log.Println(urs)
	updateRecord := make(map[string]interface{})
	updateRecord["record_episode_count"] = urs.RecordEpisodeCount - 1
	if urs.RecordEpisodeCount-1 == 0 {
		log.Println("0일 때")
		tx.Model(&urs).Updates(updateRecord)
	} else {
		log.Println("0이 아닐 때")
		tx.Updates(&urs)
	}

	return nil
}
