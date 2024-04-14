package entity

type SeriesPublisher struct {
	SeriesId    uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	PublisherId uint `gorm:"primaryKey;column:publisher_id;type:int(11);not null"`
}
