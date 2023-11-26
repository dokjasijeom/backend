package entity

import "gorm.io/gorm"

type Role struct {
	Id   uint   `gorm:"primaryKey;column:id;type:int(11)"`
	Role string `gorm:"column:role;type:varchar(10);unique;not null"`
	// Role: ADMIN, USER
}

func (Role) TableName() string {
	return "role"
}

func (Role) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (Role) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

func (Role) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}
