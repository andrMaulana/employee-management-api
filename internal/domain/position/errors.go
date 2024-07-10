package position

import "errors"

var (
	ErrPositionNotFound = errors.New("position not found")
	ErrInvalidPosition  = errors.New("invalid position data")
)
