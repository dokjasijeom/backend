package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserPasswordReset struct {
	Id        uint      `gorm:"primaryKey;column:id;type:int;not null;autoIncrement" json:"-"`
	Email     string    `gorm:"column:email;type:varchar(255);not null" json:"email,omitempty" validate:"required,email"`
	Token     string    `gorm:"column:token;type:varchar(255);not null" json:"token,omitempty"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:datetime;not null" json:"expiredAt,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
