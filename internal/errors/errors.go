package errors

import (
	"errors"
	"fmt"
	kErrors "github.com/go-kratos/kratos/v2/errors"
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
	if se := new(kErrors.Error); kErrors.As(err, &se) {
		return NewHttpError(int(se.Code), se.Reason, se.Message)
	}
	return &HttpError{Code: netHttp.StatusInternalServerError}
}
