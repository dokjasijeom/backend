package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id              uint           `gorm:"primaryKey;column:id;type:int;not null;autoIncrement" json:"-"`
	HashId          string         `gorm:"column:hash_id;type:varchar(255);null;unique" json:"hashId,omitempty"`
	Email           string         `gorm:"column:email;unique;type:varchar(255);not null" json:"email,omitempty" validate:"required,email"`
	Password        string         `gorm:"column:password;type:varchar(255);not null" json:"-"`
	Roles           []*Role        `gorm:"many2many:user_roles;"`                         // many to many relationship
	LikeSeries      []*Series      `gorm:"many2many:user_like_series;" json:"likeSeries"` // many to many relationship
	LikeSeriesCount uint           `json:"likeSeriesCount"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (User) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (User) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}
