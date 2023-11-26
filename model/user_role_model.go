package model

import "github.com/google/uuid"

type UserRoleModel struct {
	Id     uuid.UUID `json:"id"`
	Role   string    `json:"role"`
	UserId string    `json:"user_id"`
}
