package helper

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type HTTPError struct {
	Detail string
	Code   int
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Detail)
}

func NewHTTPError(detail string, code int) *HTTPError {
	return &HTTPError{
		Detail: detail,
		Code:   code,
	}
}

func (helper *ServerHelper) HandleHTTPError(err error, w http.ResponseWriter) {
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		helper.HTTPResponse(w, Error{Detail: httpErr.Detail}, httpErr.Code)
	} else if err != nil {
		log.Printf("Internal server error: %v", err)
		helper.HTTPResponse(w, Error{Detail: "Internal server error"}, http.StatusInternalServerError)
	}
}

func (helper *ServerHelper) BadRequest(detail string, args ...any) error {
	return NewHTTPError(fmt.Sprintf(detail, args...), http.StatusBadRequest)
}

func (helper *ServerHelper) Unauthorized(detail string, args ...any) error {
	return NewHTTPError(fmt.Sprintf(detail, args...), http.StatusUnauthorized)
}

func (helper *ServerHelper) NotFound(detail string, args ...any) error {
	return NewHTTPError(fmt.Sprintf(detail, args...), http.StatusNotFound)
}

func (helper *ServerHelper) Forbidden(detail string, args ...any) error {
	return NewHTTPError(fmt.Sprintf(detail, args...), http.StatusForbidden)
}

func (helper *ServerHelper) Conflict(detail string, args ...any) error {
	return NewHTTPError(fmt.Sprintf(detail, args...), http.StatusConflict)
}

func (helper *ServerHelper) UnprocessableEntity(detail string, args ...any) error {
	return NewHTTPError(fmt.Sprintf(detail, args...), http.StatusUnprocessableEntity)
}
