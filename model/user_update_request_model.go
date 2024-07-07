package model

import "os"

type UserUpdateRequestModel struct {
	Username        string  `form:"username,omitempty" validate:"required"`
	Image           os.File `form:"image,omitempty"`
	Avatar          string  `form:"avatar,omitempty"`
	Password        string  `form:"password,omitempty" validate:"required"`
	PasswordConfirm string  `form:"passwordConfirm,omitempty" validate:"required"`
}
