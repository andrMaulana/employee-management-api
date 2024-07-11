package services

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/attendance"
	"github.com/andrMaulana/employee-management-api/internal/domain/department"
	"github.com/andrMaulana/employee-management-api/internal/domain/employee"
	"github.com/andrMaulana/employee-management-api/internal/domain/location"
	"github.com/andrMaulana/employee-management-api/internal/domain/position"
)

type ReportService interface {
	GetAttendanceReport(ctx context.Context, startDate, endDate time.Time) ([]AttendanceReportItem, error)
}

type reportService struct {
	attendanceRepo attendance.Repository
	employeeRepo   employee.Repository
	departmentRepo department.Repository
	positionRepo   position.Repository
	locationRepo   location.Repository
}

type AttendanceReportItem struct {
	Date           string    `json:"date"`
	EmployeeCode   string    `json:"employee_code"`
	EmployeeName   string    `json:"employee_name"`
	DepartmentName string    `json:"department_name"`
	PositionName   string    `json:"position_name"`
	LocationName   string    `json:"location_name"`
	AbsentIn       time.Time `json:"absent_in"`
	AbsentOut      time.Time `json:"absent_out"`
}

func NewReportService(
	attendanceRepo attendance.Repository,
	employeeRepo employee.Repository,
	departmentRepo department.Repository,
	positionRepo position.Repository,
	locationRepo location.Repository,
) ReportService {
	return &reportService{
		attendanceRepo: attendanceRepo,
		employeeRepo:   employeeRepo,
		departmentRepo: departmentRepo,
		positionRepo:   positionRepo,
		locationRepo:   locationRepo,
	}
}

func (s *reportService) GetAttendanceReport(ctx context.Context, startDate, endDate time.Time) ([]AttendanceReportItem, error) {
	attendances, err := s.attendanceRepo.GetByDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var report []AttendanceReportItem

	for _, att := range attendances {
		emp, err := s.employeeRepo.GetByID(ctx, att.EmployeeID)
		if err != nil {
			continue
		}

		dept, err := s.departmentRepo.GetByID(ctx, emp.DepartmentID)
		if err != nil {
			continue
		}

		pos, err := s.positionRepo.GetByID(ctx, emp.PositionID)
		if err != nil {
			continue
		}

		loc, err := s.locationRepo.GetByID(ctx, att.LocationID)
		if err != nil {
			continue
		}

		item := AttendanceReportItem{
			Date:           att.CreatedAt.Format("2006-01-02"),
			EmployeeCode:   emp.EmployeeCode,
			EmployeeName:   emp.EmployeeName,
			DepartmentName: dept.DepartmentName,
			PositionName:   pos.PositionName,
			LocationName:   loc.LocationName,
			AbsentIn:       *att.AbsentIn,
			AbsentOut:      *att.AbsentOut,
		}

		report = append(report, item)
	}

	return report, nil
}
