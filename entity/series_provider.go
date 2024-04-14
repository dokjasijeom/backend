package entity

type SeriesProvider struct {
	SeriesId   uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	ProviderId uint `gorm:"primaryKey;column:provider_id;type:int(11);not null"`
}
