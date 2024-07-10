package department

import "time"

type Department struct {
	ID        uint       `json:"id" gorm:"primaryKey;column:Department_id"`
	Name      string     `json:"name" gorm:"column:Department_name"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:Created_at"`
	CreatedBy string     `json:"created_by" gorm:"column:Created_by"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:Updated_at"`
	UpdatedBy string     `json:"updated_by" gorm:"column:Updated_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:Deleted_at"`
}

type Repository interface {
	Create(department *Department) error
	FindAll() ([]Department, error)
	FindByID(id uint) (*Department, error)
	Update(department *Department) error
	Delete(id uint) error
}

type Service interface {
	Create(department *Department) error
	GetAll() ([]Department, error)
	GetByID(id uint) (*Department, error)
	Update(department *Department) error
	Delete(id uint) error
}
