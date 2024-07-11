package attendance

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, attendance *Attendance) error
	GetAll(ctx context.Context) ([]Attendance, error)
	GetByID(ctx context.Context, id uint) (*Attendance, error)
	Update(ctx context.Context, attendance *Attendance) error
	Delete(ctx context.Context, id uint) error
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Attendance, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, attendance *Attendance) error {
	return r.db.WithContext(ctx).Create(attendance).Error
}

func (r *repository) GetAll(ctx context.Context) ([]Attendance, error) {
	var attendances []Attendance
	err := r.db.WithContext(ctx).Find(&attendances).Error
	return attendances, err
}

func (r *repository) GetByID(ctx context.Context, id uint) (*Attendance, error) {
	var attendance Attendance
	err := r.db.WithContext(ctx).First(&attendance, id).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *repository) Update(ctx context.Context, attendance *Attendance) error {
	return r.db.WithContext(ctx).Save(attendance).Error
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&Attendance{}, id).Error
}

func (r *repository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Attendance, error) {
	var attendances []Attendance
	err := r.db.WithContext(ctx).Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&attendances).Error
	return attendances, err
}
