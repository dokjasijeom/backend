package entity

type SeriesPerson struct {
	SeriesId   uint   `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	PersonId   uint   `gorm:"primaryKey;column:person_id;type:int(11);not null"`
	PersonType string `gorm:"column:person_type;type:enum('author','illustrator','based on');not null"`
}

func (SeriesPerson) TableName() string {
	return "series_person"
}
