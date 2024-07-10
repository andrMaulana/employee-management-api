package department

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID        uint       `json:"id" gorm:"primaryKey;column:department_id"`
	Name      string     `json:"name" gorm:"column:department_name"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	CreatedBy string     `json:"created_by" gorm:"column:created_by"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy string     `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (Department) TableName() string {
	return "departments"
}

func (Department) Scopes() map[string]func(*gorm.DB) *gorm.DB {
	return map[string]func(*gorm.DB) *gorm.DB{
		"active": func(db *gorm.DB) *gorm.DB {
			return db.Where("Deleted_at IS NULL")
		},
	}
}

// type Repository interface {
// 	Create(department *Department) error
// 	FindAll() ([]Department, error)
// 	FindByID(id uint) (*Department, error)
// 	Update(department *Department) error
// 	Delete(id uint) error
// }

// type Service interface {
// 	Create(department *Department) error
// 	GetAll() ([]Department, error)
// 	GetByID(id uint) (*Department, error)
// 	Update(department *Department) error
// 	Delete(id uint) error
// }
