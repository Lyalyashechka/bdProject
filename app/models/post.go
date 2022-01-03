package models

type Post struct {
	Id int64 `json:"id"`
	Parent int64 `json:"parent"`
	Author string `json:"author" validate:"required"`
	Message string `json:"message" validate:"required"`
	IsEdited bool `json:"is_edited"`
	Forum string `json:"forum"`
	Thread int32 `json:"thread"`
	Created string `json:"created"`
}