package entity

type SeriesAuthor struct {
	SeriesId   uint   `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	PersonId   uint   `gorm:"primaryKey;column:person_id;type:int(11);not null"`
	PersonType string `gorm:"column:person_type;type:varchar(255);null;index;" json:"personType,omitempty"`
	// person_type: author, illustrator, original_author
}

func (SeriesAuthor) TableName() string {
	return "series_authors"
}
