package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserProfile struct {
	Id        uint           `gorm:"primaryKey;column:id;type:int;not null;autoIncrement" json:"-"`
	UserId    uint           `gorm:"column:user_id;type:int;not null" json:"-"`
	Username  string         `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Avatar    string         `gorm:"column:avatar;type:varchar(255);null" json:"avatar"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
