package services

import (
	"context"
	"log"

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
	// emp, err := s.employeeRepo.GetByEmployeeCode(ctx, employeeCode)
	// if err != nil {
	// 	return "", errors.ErrInvalidCredentials
	// }

	// if err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(password)); err != nil {
	// 	return "", errors.ErrInvalidCredentials
	// }

	// token, err := auth.GenerateToken(emp.EmployeeID, emp.EmployeeName)
	// if err != nil {
	// 	return "", errors.ErrInternalServer
	// }

	// return token, nil

	log.Printf("Attempting login for employee code: %s", employeeCode)

	emp, err := s.employeeRepo.GetByEmployeeCode(ctx, employeeCode)
	if err != nil {
		if err == errors.ErrNotFound {
			log.Printf("Login failed: Employee with code %s not found", employeeCode)
			return "", errors.ErrInvalidCredentials
		}
		log.Printf("Login failed: Error fetching employee data: %v", err)
		return "", errors.ErrInternalServer
	}

	if err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(password)); err != nil {
		log.Printf("Login failed: Invalid password for employee code %s", employeeCode)
		return "", errors.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(emp.EmployeeID, emp.EmployeeName)
	if err != nil {
		log.Printf("Login failed: Error generating token: %v", err)
		return "", errors.ErrInternalServer
	}

	log.Printf("Login successful for employee code: %s", employeeCode)
	return token, nil
}
