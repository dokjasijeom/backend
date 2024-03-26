package entity

// insert row after user like series than series like count increase
type UserLikeSeriesCount struct {
	SeriesId uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	Count    uint `gorm:"column:count;type:int(11);not null"`
}

func (UserLikeSeriesCount) TableName() string {
	return "user_like_series_count"
}
