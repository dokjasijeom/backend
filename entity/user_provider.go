package entity

type UserProvider struct {
	UserId     uint `gorm:"primaryKey;column:user_id;type:int(11);not null" json:"userId,omitempty"`
	ProviderId uint `gorm:"primaryKey;column:provider_id;type:int(11);not null" json:"providerId,omitempty"`
}
