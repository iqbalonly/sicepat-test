package ierr

import (
	"errors"
	"net/http"
	"sicepat/internal/constant"
)

var (
	FailedParameterType = &Error{
		Code: constant.BadRequestGeneral,
		Err:  errors.New("failed parameter type"),
	}

	InvalidRequiredParameter = &Error{
		Code: constant.BadRequestGeneral,
		Err:  errors.New("invalid required parameter"),
	}

	ErrInvalidDate = &Error{
		Code: constant.BadRequestGeneral,
		Err:  errors.New("invalid date format, acceptable yyyy-dd-mm"),
	}

	NoRowsAffected = &Error{
		Code: constant.InternalServerError,
		Err:  errors.New("no rows affected, please check your request"),
	}

	ErrUnknown = &Error{
		Code: constant.Unknown,
		Err:  errors.New("unknown error"),
	}
)

//New creates new internal error wrapper
func New(code constant.Code, err error) *Error {
	return &Error{
		code, err,
	}
}

type IerrInterface interface {
	GetCode() constant.Code
	GetHTTPCode() int
	Error() string
	Unwrap() error
}

type Error struct {
	Code constant.Code `json:"code"`
	Err  error         `json:"err"`
}

func (e *Error) GetCode() constant.Code {
	return e.Code
}

func (e *Error) GetHTTPCode() int {
	var statusCode int

	switch e.GetCode() {
	case constant.Success:
		statusCode = http.StatusOK

	case constant.BadRequestGeneral:
		statusCode = http.StatusBadRequest

	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Error() string {
	return e.Err.Error()
}
