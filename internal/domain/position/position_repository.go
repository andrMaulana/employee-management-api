package position

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, position *Position) error
	GetAll(ctx context.Context) ([]Position, error)
	GetByID(ctx context.Context, id uint) (*Position, error)
	Update(ctx context.Context, position *Position) error
	Delete(ctx context.Context, id uint) error
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

func (r *repository) GetAll(ctx context.Context) ([]Position, error) {
	var positions []Position
	err := r.db.WithContext(ctx).Find(&positions).Error
	return positions, err
}

func (r *repository) GetByID(ctx context.Context, id uint) (*Position, error) {
	var position Position
	err := r.db.WithContext(ctx).First(&position, id).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *repository) Update(ctx context.Context, position *Position) error {
	return r.db.WithContext(ctx).Save(position).Error
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Position{}, id).Error
}
