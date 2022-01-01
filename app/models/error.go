package models

const (
	NoUser       = "Can't find user\n"
	ConflictData = "Conflict data\n"
	NoSlug       = "Can't find slug\n"
)

const (
	PgxUniqErrorCode         = "23505"
	PgxNoFoundFieldErrorCode = "23503"
)

type CustomError struct {
	Message string `json:"message"`
}
