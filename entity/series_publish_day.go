package entity

type SeriesPublishDay struct {
	SeriesId     uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	PublishDayId uint `gorm:"primaryKey;column:publish_day_id;type:int(11);not null"`
}

func (SeriesPublishDay) TableName() string {
	return "series_publish_days"
}
