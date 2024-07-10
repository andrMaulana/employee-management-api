package attendance

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID         uint       `json:"id" gorm:"primaryKey;column:attendance_id"`
	EmployeeID uint       `json:"employee_id" gorm:"column:employee_id"`
	LocationID uint       `json:"location_id" gorm:"column:location_id"`
	AbsentIn   time.Time  `json:"absent_in" gorm:"column:absent_in"`
	AbsentOut  *time.Time `json:"absent_out,omitempty" gorm:"column:absent_out"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at"`
	CreatedBy  string     `json:"created_by" gorm:"column:created_by"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy  string     `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (Attendance) TableName() string {
	return "attendance"
}

func (Attendance) Scopes() map[string]func(*gorm.DB) *gorm.DB {
	return map[string]func(*gorm.DB) *gorm.DB{
		"active": func(db *gorm.DB) *gorm.DB {
			return db.Where("Deleted_at IS NULL")
		},
	}
}
