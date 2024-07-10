package department

import (
	"context"
	"errors"
	"time"

	"github.com/andrMaulana/employee-management-api/pkg/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, department *Department) error
	FindAll(ctx context.Context, params *FindAllParams) ([]Department, int64, error)
	FindByID(ctx context.Context, id uint) (*Department, error)
	Update(ctx context.Context, department *Department) error
	Delete(ctx context.Context, id uint) error
	BatchCreate(ctx context.Context, departments []Department) error
	BatchUpdate(ctx context.Context, departments []Department) error
	BatchDelete(ctx context.Context, ids []uint) error
}

type FindAllParams struct {
	pagination.Paginator
	Search    string
	SortBy    string
	SortOrder string
	CreatedAt *time.Time
	UpdatedAt *time.Time
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

func (r *repository) FindAll(ctx context.Context, params *FindAllParams) ([]Department, int64, error) {
	var departments []Department
	var count int64

	query := r.db.WithContext(ctx).Model(&Department{}).Scopes(Department{}.Scopes()["active"])

	if params.Search != "" {
		query = query.Where("Department_name ILIKE ?", "%"+params.Search+"%")
	}

	if params.CreatedAt != nil {
		query = query.Where("Created_at >= ?", params.CreatedAt)
	}

	if params.UpdatedAt != nil {
		query = query.Where("Updated_at >= ?", params.UpdatedAt)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if params.SortBy != "" {
		order := params.SortBy
		if params.SortOrder == "desc" {
			order += " DESC"
		}
		query = query.Order(order)
	}

	err := query.Offset(params.GetOffset()).Limit(params.GetLimit()).Find(&departments).Error
	return departments, count, err
}

func (r *repository) FindByID(ctx context.Context, id uint) (*Department, error) {
	var department Department
	err := r.db.WithContext(ctx).Scopes(Department{}.Scopes()["active"]).Where("Department_id = ?", id).First(&department).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDepartmentNotFound
		}
		return nil, err
	}
	return &department, nil
}

func (r *repository) Update(ctx context.Context, department *Department) error {
	result := r.db.WithContext(ctx).Model(department).Updates(Department{
		Name:      department.Name,
		UpdatedAt: time.Now(),
		UpdatedBy: department.UpdatedBy,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrDepartmentNotFound
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&Department{}).Where("Department_id = ?", id).Update("Deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrDepartmentNotFound
	}
	return nil
}

func (r *repository) BatchCreate(ctx context.Context, departments []Department) error {
	return r.db.WithContext(ctx).Create(&departments).Error
}

func (r *repository) BatchUpdate(ctx context.Context, departments []Department) error {
	for _, dept := range departments {
		if err := r.Update(ctx, &dept); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) BatchDelete(ctx context.Context, ids []uint) error {
	return r.db.WithContext(ctx).Model(&Department{}).Where("Department_id IN ?", ids).Update("Deleted_at", time.Now()).Error
}
