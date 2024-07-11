package position

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	PositionID   uint           `gorm:"primaryKey" json:"position_id"`
	DepartmentID uint           `json:"department_id"`
	PositionName string         `gorm:"size:255" json:"position_name"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `gorm:"size:255" json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    string         `gorm:"size:255" json:"updated_by"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
