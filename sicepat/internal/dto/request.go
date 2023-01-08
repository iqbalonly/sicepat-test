package dto

import (
	"sicepat/internal/ierr"
	"time"
)

const (
	RequestDateLayout = "2006-01-02"
)

type RequestAble interface {
	Validate() error
}

type UserRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}

func (r *UserRequest) Validate() error {
	if r.Name == "" || r.Email == "" {
		return ierr.InvalidRequiredParameter
	}

	if r.DateOfBirth != "" {
		_, err := time.Parse(RequestDateLayout, r.DateOfBirth)
		if err != nil {
			return ierr.ErrInvalidDate
		}
	}

	return nil
}
