package department

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, department *Department) error
	GetAll(ctx context.Context) ([]Department, error)
	GetByID(ctx context.Context, id uint) (*Department, error)
	Update(ctx context.Context, department *Department) error
	Delete(ctx context.Context, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, department *Department) error {
	return r.db.WithContext(ctx).Create(department).Error
}

func (r *repository) GetAll(ctx context.Context) ([]Department, error) {
	var departments []Department
	err := r.db.WithContext(ctx).Find(&departments).Error
	return departments, err
}

func (r *repository) GetByID(ctx context.Context, id uint) (*Department, error) {
	var department Department
	err := r.db.WithContext(ctx).First(&department, id).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *repository) Update(ctx context.Context, department *Department) error {
	return r.db.WithContext(ctx).Save(department).Error
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Department{}, id).Error
}
