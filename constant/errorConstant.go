package constant

type ErrorMessage string

const (
	ErrBadRequest          ErrorMessage = "error bad request"
	ErrInternalServerError ErrorMessage = "internal server error"
	ErrUnauthorized        ErrorMessage = "unauthorized"
	ErrForbidden           ErrorMessage = "forbidden"
	ErrNotFound            ErrorMessage = "not found"
)
