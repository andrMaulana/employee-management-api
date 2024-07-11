package employee

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeID   uint           `gorm:"primaryKey" json:"employee_id"`
	EmployeeCode string         `gorm:"size:255;unique" json:"employee_code"`
	EmployeeName string         `gorm:"size:255" json:"employee_name"`
	Password     string         `gorm:"size:255" json:"-"`
	DepartmentID uint           `json:"department_id"`
	PositionID   uint           `json:"position_id"`
	Superior     uint           `json:"superior"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `gorm:"size:255" json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    string         `gorm:"size:255" json:"updated_by"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
