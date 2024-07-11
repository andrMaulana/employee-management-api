package services

import (
	"context"
	"fmt"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/employee"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, name, password string, departmentID, positionID, superior uint, createdBy string) (*employee.Employee, error)
	GetAllEmployees(ctx context.Context) ([]employee.Employee, error)
	GetEmployeeByID(ctx context.Context, id uint) (*employee.Employee, error)
	UpdateEmployee(ctx context.Context, id uint, name string, departmentID, positionID, superior uint, updatedBy string) (*employee.Employee, error)
	DeleteEmployee(ctx context.Context, id uint) error
}

type employeeService struct {
	repo employee.Repository
}

func NewEmployeeService(repo employee.Repository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(ctx context.Context, name, password string, departmentID, positionID, superior uint, createdBy string) (*employee.Employee, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	employeeCode := generateEmployeeCode()

	emp := &employee.Employee{
		EmployeeCode: employeeCode,
		EmployeeName: name,
		Password:     string(hashedPassword),
		DepartmentID: departmentID,
		PositionID:   positionID,
		Superior:     superior,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, emp); err != nil {
		return nil, errors.ErrInternalServer
	}

	return emp, nil
}

func (s *employeeService) GetAllEmployees(ctx context.Context) ([]employee.Employee, error) {
	return s.repo.GetAll(ctx)
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, id uint) (*employee.Employee, error) {
	emp, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return emp, nil
}

func (s *employeeService) UpdateEmployee(ctx context.Context, id uint, name string, departmentID, positionID, superior uint, updatedBy string) (*employee.Employee, error) {
	emp, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	emp.EmployeeName = name
	emp.DepartmentID = departmentID
	emp.PositionID = positionID
	emp.Superior = superior
	emp.UpdatedBy = updatedBy
	emp.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, emp); err != nil {
		return nil, errors.ErrInternalServer
	}

	return emp, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.ErrInternalServer
	}
	return nil
}

func generateEmployeeCode() string {
	now := time.Now()
	return fmt.Sprintf("%02d%02d%04d", now.Year()%100, now.Month(), now.Unix()%10000)
}
