package location

import (
	"context"
	"errors"
	"time"

	"github.com/andrMaulana/employee-management-api/pkg/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, location *Location) error
	FindAll(ctx context.Context, params *FindAllParams) ([]Location, int64, error)
	FindByID(ctx context.Context, id uint) (*Location, error)
	Update(ctx context.Context, location *Location) error
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

func (r *repository) Create(ctx context.Context, location *Location) error {
	return r.db.WithContext(ctx).Create(location).Error
}

func (r *repository) FindAll(ctx context.Context, params *FindAllParams) ([]Location, int64, error) {
	var location []Location
	var count int64

	query := r.db.WithContext(ctx).Model(&Location{}).Scopes(Location{}.Scopes()["active"])

	if params.Search != "" {
		query = query.Where("Location_name ILIKE ?", "%"+params.Search+"%")
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

	err := query.Offset(params.GetOffset()).Limit(params.GetLimit()).Find(&location).Error
	return location, count, err
}

func (r *repository) FindByID(ctx context.Context, id uint) (*Location, error) {
	var location Location
	err := r.db.WithContext(ctx).Scopes(Location{}.Scopes()["active"]).Where("Location_id = ?", id).First(&location).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLocationNotFound
		}
		return nil, err
	}
	return &location, nil
}

func (r *repository) Update(ctx context.Context, location *Location) error {
	result := r.db.WithContext(ctx).Model(location).Updates(Location{
		Name:      location.Name,
		UpdatedAt: time.Now(),
		UpdatedBy: location.UpdatedBy,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrLocationNotFound
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&Location{}).Where("Location_id = ?", id).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrLocationNotFound
	}
	return nil
}
