package entity

import (
	"gorm.io/gorm"
	"time"
)

type SeriesType string

const (
	WebNovel SeriesType = "webnovel"
	WebToon  SeriesType = "webtoon"
)

type Series struct {
	Id          uint         `gorm:"primaryKey;column:id;type:int(11);not null;autoIncrement" json:"id,omitempty"`
	HashId      string       `gorm:"column:hash_id;type:varchar(255);null" json:"hashId"`
	Title       string       `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string       `gorm:"column:description;type:text;null" json:"description"`
	Thumbnail   string       `gorm:"column:thumbnail;type:varchar(255);not null" json:"thumbnail"`
	ISBN        string       `gorm:"column:isbn;type:varchar(255);null" json:"isbn"`
	ECN         string       `gorm:"column:ecn;type:varchar(255);null" json:"ecn"`
	SeriesType  SeriesType   `gorm:"column:series_type;type:varchar(20);not null" json:"seriesType"`
	PublishDays []PublishDay `gorm:"many2many:series_publish_days;" json:"publishDays"` // many to many relationship
	Episodes    []Episode    `gorm:"many2many:series_episodes;" json:"episodes"`        // many to many relationship
	//Persons     []Person     `gorm:"many2many:series_persons;" json:"people"`           // many to many relationship
	Genres        []Genre        `gorm:"many2many:series_genres;" json:"genres"` // many to many relationship
	SeriesAuthors []SeriesAuthor `json:"-"`                                      // many to many relationship
	Authors       []Person       `gorm:"-" json:"authors"`
	Publishers    []Publisher    `gorm:"many2many:series_publishers;" json:"publishers"` // many to many relationship
	Providers     []Provider     `gorm:"many2many:series_providers;" json:"providers"`   // many to many relationship
	LikeCount     uint           `gorm:"column:like_count;type:int(11);not null;default:0;index" json:"likeCount"`
	DisplayTags   string         `gorm:"-" json:"displayTags"`
	TotalEpisode  uint           `gorm:"-" json:"totalEpisode"`
	IsComplete    bool           `gorm:"column:is_complete;type:tinyint(1);not null" json:"isComplete,default:false"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Series) TableName() string {
	return "series"
}

// after delete
func (s *Series) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&SeriesEpisode{}).Where("series_id = ?", s.Id).Delete(&SeriesEpisode{})
	tx.Model(&SeriesGenre{}).Where("series_id = ?", s.Id).Delete(&SeriesGenre{})
	tx.Model(&SeriesAuthor{}).Where("series_id = ?", s.Id).Delete(&SeriesAuthor{})
	tx.Model(&SeriesPublisher{}).Where("series_id = ?", s.Id).Delete(&SeriesPublisher{})
	tx.Model(&SeriesProvider{}).Where("series_id = ?", s.Id).Delete(&SeriesProvider{})
	tx.Model(&SeriesPublishDay{}).Where("series_id = ?", s.Id).Delete(&SeriesPublishDay{})
	return
}
