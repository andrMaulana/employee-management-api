package location

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, location *Location) error
	GetAll(ctx context.Context) ([]Location, error)
	GetByID(ctx context.Context, id uint) (*Location, error)
	Update(ctx context.Context, location *Location) error
	Delete(ctx context.Context, id uint) error
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

func (r *repository) GetAll(ctx context.Context) ([]Location, error) {
	var locations []Location
	err := r.db.WithContext(ctx).Find(&locations).Error
	return locations, err
}

func (r *repository) GetByID(ctx context.Context, id uint) (*Location, error) {
	var location Location
	err := r.db.WithContext(ctx).First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *repository) Update(ctx context.Context, location *Location) error {
	return r.db.WithContext(ctx).Save(location).Error
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Location{}, id).Error
}
