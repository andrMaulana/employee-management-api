package employee

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID           uint       `json:"id" gorm:"primaryKey;column:employee_id"`
	Code         string     `json:"code" gorm:"column:employee_code"`
	Name         string     `json:"name" gorm:"column:employee_name"`
	Password     string     `json:"password" gorm:"column:password"`
	DepartmentID uint       `json:"department_id" gorm:"column:department_id"`
	PositionID   uint       `json:"position_id" gorm:"column:position_id"`
	Superior     uint       `json:"superior" gorm:"column:superior"`
	CreatedAt    time.Time  `json:"created_at" gorm:"column:created_at"`
	CreatedBy    string     `json:"created_by" gorm:"column:created_by"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy    string     `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (Employee) TableName() string {
	return "employee"
}

func (Employee) Scopes() map[string]func(*gorm.DB) *gorm.DB {
	return map[string]func(*gorm.DB) *gorm.DB{
		"active": func(db *gorm.DB) *gorm.DB {
			return db.Where("Deleted_at IS NULL")
		},
	}
}
