package models

type Forum struct {
	Title   string `json:"title" validate:"required"`
	User    string `json:"user" validate:"required"`
	Slug    string `json:"slug" validate:"required"`
	Posts   int64  `json:"posts,omitempty"`
	Threads int64  `json:"threads,omitempty"`
}
