package errors

import "net/http"

type AppError struct {
	Id       string
	HttpCode int
}

func (e *AppError) Error() string {
	return e.Id
}

var (
	InternalError = &AppError{"Internal error", http.StatusInternalServerError}
	Unauthorized  = &AppError{"Unauthorized", http.StatusUnauthorized}
	BadRequest    = &AppError{"Bad request", http.StatusBadRequest}
	InvalidToken  = &AppError{"Invalid token", http.StatusUnauthorized}
)
