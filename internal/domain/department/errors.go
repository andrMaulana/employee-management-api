// internal/domain/department/errors.go

package department

import "errors"

var (
	ErrDepartmentNotFound = errors.New("department not found")
	ErrInvalidDepartment  = errors.New("invalid department data")
)
