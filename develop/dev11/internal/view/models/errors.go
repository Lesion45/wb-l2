package models

import "errors"

var ErrInvalidData error

func NewErrInvalidData(msg string) error {
	ErrInvalidData = errors.New("invalid data: " + msg)
	return ErrInvalidData
}
