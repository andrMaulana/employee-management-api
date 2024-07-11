package services

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/department"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
)

type DepartmentService interface {
	CreateDepartment(ctx context.Context, name, createdBy string) (*department.Department, error)
	GetAllDepartments(ctx context.Context) ([]department.Department, error)
	GetDepartmentByID(ctx context.Context, id uint) (*department.Department, error)
	UpdateDepartment(ctx context.Context, id uint, name, updatedBy string) (*department.Department, error)
	DeleteDepartment(ctx context.Context, id uint) error
}

type departmentService struct {
	repo department.Repository
}

func NewDepartmentService(repo department.Repository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) CreateDepartment(ctx context.Context, name, createdBy string) (*department.Department, error) {
	dept := &department.Department{
		DepartmentName: name,
		CreatedBy:      createdBy,
		UpdatedBy:      createdBy,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.repo.Create(ctx, dept); err != nil {
		return nil, errors.ErrInternalServer
	}

	return dept, nil
}

func (s *departmentService) GetAllDepartments(ctx context.Context) ([]department.Department, error) {
	return s.repo.GetAll(ctx)
}

func (s *departmentService) GetDepartmentByID(ctx context.Context, id uint) (*department.Department, error) {
	dept, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return dept, nil
}

func (s *departmentService) UpdateDepartment(ctx context.Context, id uint, name, updatedBy string) (*department.Department, error) {
	dept, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	dept.DepartmentName = name
	dept.UpdatedBy = updatedBy
	dept.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, dept); err != nil {
		return nil, errors.ErrInternalServer
	}

	return dept, nil
}

func (s *departmentService) DeleteDepartment(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.ErrInternalServer
	}
	return nil
}
