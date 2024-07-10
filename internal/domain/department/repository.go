package department

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(department *Department) error {
	return r.db.Create(department).Error
}

func (r *repository) FindAll() ([]Department, error) {
	var departments []Department
	err := r.db.Where("Deleted_at IS NULL").Find(&departments).Error
	return departments, err
}

func (r *repository) FindByID(id uint) (*Department, error) {
	var department Department
	err := r.db.Where("Department_id = ? AND Deleted_at IS NULL", id).First(&department).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *repository) Update(department *Department) error {
	return r.db.Save(department).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Model(&Department{}).Where("Department_id = ?", id).Update("Deleted_at", gorm.Expr("NOW()")).Error
}
