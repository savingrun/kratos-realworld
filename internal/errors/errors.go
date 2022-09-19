package errors

import (
	"errors"
	"fmt"
	netHttp "net/http"
)

func NewHttpError(code int, field, detail string) *HttpError {
	return &HttpError{
		Code: code,
		Errors: map[string][]string{
			field: {detail},
		},
	}
}

type HttpError struct {
	Errors map[string][]string `json:"errors"`

	Code int `json:"-"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("HttpError: %d", e.Code)
}

func FromError(err error) *HttpError {
	if err == nil {
		return nil
	}
	if se := new(HttpError); errors.As(err, &se) {
		return se
	}
	return &HttpError{Code: netHttp.StatusInternalServerError}
}
