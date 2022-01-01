package models

type User struct {
	Nickname string `json:"nickname,omitempty"`
	FullName string `json:"fullname" validate:"required"`
	About    string `json:"about"`
	Email    string `json:"email"    validate:"required,email"`
}

type UserUpdate struct {
	FullName string `json:"fullname" validate:"required"`
	About    string `json:"about" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
}
