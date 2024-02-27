package entity

type SeriesAuthor struct {
	SeriesId uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	PersonId uint `gorm:"primaryKey;column:person_id;type:int(11);not null"`
}

func (SeriesAuthor) TableName() string {
	return "series_authors"
}
