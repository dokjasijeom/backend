package entity

type SeriesGenre struct {
	SeriesId uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	GenreId  uint `gorm:"primaryKey;column:genre_id;type:int(11);not null"`
}

func (SeriesGenre) TableName() string {
	return "series_genres"
}
