package entity

import "gorm.io/gorm"

type Role struct {
	Id    uint    `gorm:"primaryKey;column:id;type:int(11)" json:"-"`
	Role  string  `gorm:"column:role;type:varchar(10);unique;not null"`
	Users []*User `gorm:"many2many:user_roles" json:"-"`
	// Role: ADMIN, USER
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
