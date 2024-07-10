package position

import (
	"context"
	"errors"
	"time"

	"github.com/andrMaulana/employee-management-api/pkg/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, position *Position) error
	FindAll(ctx context.Context, params *FindAllParams) ([]Position, int64, error)
	FindByID(ctx context.Context, id uint) (*Position, error)
	Update(ctx context.Context, position *Position) error
	Delete(ctx context.Context, id uint) error
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

func (r *repository) Create(ctx context.Context, position *Position) error {
	return r.db.WithContext(ctx).Create(position).Error
}

func (r *repository) FindAll(ctx context.Context, params *FindAllParams) ([]Position, int64, error) {
	var positions []Position
	var count int64

	query := r.db.WithContext(ctx).Model(&Position{}).Scopes(Position{}.Scopes()["active"])

	if params.Search != "" {
		query = query.Where("Position_name ILIKE ?", "%"+params.Search+"%")
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

	err := query.Offset(params.GetOffset()).Limit(params.GetLimit()).Find(&positions).Error
	return positions, count, err
}

func (r *repository) FindByID(ctx context.Context, id uint) (*Position, error) {
	var position Position
	err := r.db.WithContext(ctx).Scopes(Position{}.Scopes()["active"]).Where("Position_id = ?", id).First(&position).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPositionNotFound
		}
		return nil, err
	}
	return &position, nil
}

func (r *repository) Update(ctx context.Context, position *Position) error {
	result := r.db.WithContext(ctx).Model(position).Updates(Position{
		Name:      position.Name,
		UpdatedAt: time.Now(),
		UpdatedBy: position.UpdatedBy,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPositionNotFound
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&Position{}).Where("Position_id = ?", id).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPositionNotFound
	}
	return nil
}
