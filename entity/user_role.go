package entity

import (
	"gorm.io/gorm"
)

type UserRole struct {
	UserId uint `gorm:"column:user_id;type:int(11);not null"`
	RoleId uint `gorm:"column:role_id;type:int(11);not null"`
}

func (UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (UserRole) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (UserRole) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}
