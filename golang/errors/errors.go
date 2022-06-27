package errors

import (
	"errors"
	"net/http"
)

var (
	BadRequestMethod = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	InternalError    = errors.New(http.StatusText(http.StatusInternalServerError))

	NoJSONBody = errors.New("Unable to decode JSON")

	MintsNotFound = errors.New("Mints not queried for block")

	BadCSRF           = errors.New("Missing CSRF Header")
	BadOrigin         = errors.New("Invalid Origin Header")
	RouteUnauthorized = errors.New("You don't have permission to view this resource")
	RouteNotFound     = errors.New("Route not found")
	ExpiredToken      = errors.New("Your access token expired")
	InvalidToken      = errors.New("Your access token is invalid")
)

func codeMap() map[error]int {
	return map[error]int{
		BadRequestMethod: http.StatusMethodNotAllowed,
		InternalError:    http.StatusInternalServerError,

		NoJSONBody: http.StatusBadRequest,

		MintsNotFound: http.StatusNotFound,

		BadCSRF:           http.StatusUnauthorized,
		BadOrigin:         http.StatusUnauthorized,
		RouteUnauthorized: http.StatusUnauthorized,
		RouteNotFound:     http.StatusNotFound,
		ExpiredToken:      http.StatusUnauthorized,
	}
}

func GetCode(e error) (bool, int) {
	if code, ok := codeMap()[e]; ok {
		return true, code
	}
	return false, http.StatusInternalServerError
}
