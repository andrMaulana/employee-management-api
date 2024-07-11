package attendance

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	AttendanceID uint           `gorm:"primaryKey" json:"attendance_id"`
	EmployeeID   uint           `gorm:"not null" json:"employee_id"`
	LocationID   uint           `gorm:"not null" json:"location_id"`
	AbsentIn     *time.Time     `json:"absent_in"`
	AbsentOut    *time.Time     `json:"absent_out"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `gorm:"size:255" json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    string         `gorm:"size:255" json:"updated_by"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
