package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	Id     uuid.UUID `gorm:"primaryKey;column:user_role_id;type:varchar(36)"`
	Roles  []Role    `gorm:"foreignKey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId uint
}

func (UserRole) TableName() string {
	return "user_role"
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
