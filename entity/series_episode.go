package entity

type SeriesEpisode struct {
	SeriesId  uint `gorm:"primaryKey;column:series_id;type:int(11);not null"`
	EpisodeId uint `gorm:"primaryKey;column:episode_id;type:int(11);not null"`
}
