package models

const (
	NoUser       = "Can't find user\n"
	ConflictData = "Conflict data\n"
)

type CustomError struct {
	Message string `json:"message"`
}
