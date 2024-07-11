package location

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	LocationID   uint           `gorm:"primaryKey" json:"location_id"`
	LocationName string         `gorm:"size:255" json:"location_name"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `gorm:"size:255" json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    string         `gorm:"size:255" json:"updated_by"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
