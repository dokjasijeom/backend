package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint   `gorm:"primaryKey;column:id;type:int;not null;primaryKey;autoIncrement"`
	HashId    string `gorm:"column:hash_id;type:varchar(255);not null"`
	Email     string `gorm:"column:email;type:varchar(255);not null"`
	Password  string `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "user"
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
