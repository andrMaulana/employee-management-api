package department

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	DepartmentID   uint           `gorm:"primaryKey" json:"department_id"`
	DepartmentName string         `gorm:"size:255;not null" json:"department_name"`
	CreatedAt      time.Time      `json:"created_at"`
	CreatedBy      string         `gorm:"size:255" json:"created_by"`
	UpdatedAt      time.Time      `json:"updated_at"`
	UpdatedBy      string         `gorm:"size:255" json:"updated_by"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
