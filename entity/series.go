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
	HashId      string       `gorm:"column:hash_id;type:varchar(255);null" json:"hashId,omitempty"`
	Title       string       `gorm:"column:title;type:varchar(255);not null" json:"title,omitempty"`
	Description string       `gorm:"column:description;type:text;null" json:"description,omitempty"`
	Thumbnail   string       `gorm:"column:thumbnail;type:varchar(255);not null" json:"thumbnail,omitempty"`
	ISBN        string       `gorm:"column:isbn;type:varchar(255);null" json:"isbn,omitempty"`
	ECN         string       `gorm:"column:ecn;type:varchar(255);null" json:"ecn,omitempty"`
	SeriesType  SeriesType   `gorm:"column:series_type;type:varchar(20);not null" json:"seriesType,omitempty"`
	PublishDays []PublishDay `gorm:"many2many:series_publish_days;" json:"publishDays,omitempty"` // many to many relationship
	Episodes    []Episode    `gorm:"many2many:series_episodes;" json:"episodes,omitempty"`        // many to many relationship
	//Persons     []Person     `gorm:"many2many:series_persons;" json:"people"`           // many to many relationship
	Genres         []Genre          `gorm:"many2many:series_genres;" json:"genres,omitempty"` // many to many relationship
	SeriesAuthors  []SeriesAuthor   `json:"-"`                                                // many to many relationship
	Authors        []Person         `gorm:"-" json:"authors,omitempty"`
	Publishers     []Publisher      `gorm:"many2many:series_publishers;" json:"publishers,omitempty"` // many to many relationship
	SeriesProvider []SeriesProvider `json:"-"`                                                        // many to many relationship
	Providers      []Provider       `gorm:"-" json:"providers,omitempty"`
	//Providers      []Provider       `gorm:"many2many:series_providers;" json:"providers"`   // many to many relationship
	LikeCount    uint   `gorm:"column:like_count;type:int(11);not null;default:0;index" json:"likeCount,omitempty"`
	ViewCount    uint   `gorm:"column:view_count;type:int(11);not null;default:0;index" json:"viewCount,omitempty"`
	DisplayTags  string `gorm:"-" json:"displayTags,omitempty"`
	TotalEpisode uint   `gorm:"column:total_episode;type:int(11);not null;default:0;index" json:"totalEpisode"`
	IsComplete   bool   `gorm:"column:is_complete;type:tinyint(1);not null" json:"isComplete,default:false"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
