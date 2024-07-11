package employee

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, employee *Employee) error
	GetAll(ctx context.Context) ([]Employee, error)
	GetByID(ctx context.Context, id uint) (*Employee, error)
	Update(ctx context.Context, employee *Employee) error
	Delete(ctx context.Context, id uint) error
	GetByEmployeeCode(ctx context.Context, code string) (*Employee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, employee *Employee) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

func (r *repository) GetAll(ctx context.Context) ([]Employee, error) {
	var employees []Employee
	err := r.db.WithContext(ctx).Find(&employees).Error
	return employees, err
}

func (r *repository) GetByID(ctx context.Context, id uint) (*Employee, error) {
	var employee Employee
	err := r.db.WithContext(ctx).First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *repository) Update(ctx context.Context, employee *Employee) error {
	return r.db.WithContext(ctx).Save(employee).Error
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Employee{}, id).Error
}

func (r *repository) GetByEmployeeCode(ctx context.Context, code string) (*Employee, error) {
	var employee Employee
	err := r.db.WithContext(ctx).Where("employee_code = ?", code).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
