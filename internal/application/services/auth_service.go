package services

import (
	"context"

	"github.com/andrMaulana/employee-management-api/internal/domain/employee"
	"github.com/andrMaulana/employee-management-api/internal/infrastructure/auth"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, employeeCode, password string) (string, error)
}

type authService struct {
	employeeRepo employee.Repository
}

func NewAuthService(employeeRepo employee.Repository) AuthService {
	return &authService{employeeRepo: employeeRepo}
}

func (s *authService) Login(ctx context.Context, employeeCode, password string) (string, error) {
	emp, err := s.employeeRepo.GetByEmployeeCode(ctx, employeeCode)
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(password)); err != nil {
		return "", errors.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(emp.EmployeeID, emp.EmployeeName)
	if err != nil {
		return "", errors.ErrInternalServer
	}

	return token, nil
}
