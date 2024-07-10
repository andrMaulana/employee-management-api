package location

import "errors"

var (
	ErrLocationNotFound = errors.New("location not found")
	ErrInvalidLocation  = errors.New("invalid location data")
)
