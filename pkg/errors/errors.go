package errors

import "errors"

var (
	ErrNotFound           = errors.New("resource not found")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInternalServer     = errors.New("internal server error")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrDuplicateResource  = errors.New("resource already exists")
	ErrInvalidCredentials = errors.New("token not valid")
)
