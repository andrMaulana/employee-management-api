package services

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/attendance"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
)

type AttendanceService interface {
	CreateAttendance(ctx context.Context, employeeID, locationID uint, absentIn, absentOut *time.Time, createdBy string) (*attendance.Attendance, error)
	GetAllAttendances(ctx context.Context) ([]attendance.Attendance, error)
	GetAttendanceByID(ctx context.Context, id uint) (*attendance.Attendance, error)
	UpdateAttendance(ctx context.Context, id uint, employeeID, locationID uint, absentIn, absentOut *time.Time, updatedBy string) (*attendance.Attendance, error)
	DeleteAttendance(ctx context.Context, id uint) error
	GetAttendanceReport(ctx context.Context, startDate, endDate time.Time) ([]attendance.Attendance, error)
}

type attendanceService struct {
	repo attendance.Repository
}

func NewAttendanceService(repo attendance.Repository) AttendanceService {
	return &attendanceService{repo: repo}
}

func (s *attendanceService) CreateAttendance(ctx context.Context, employeeID, locationID uint, absentIn, absentOut *time.Time, createdBy string) (*attendance.Attendance, error) {
	att := &attendance.Attendance{
		EmployeeID: employeeID,
		LocationID: locationID,
		AbsentIn:   absentIn,
		AbsentOut:  absentOut,
		CreatedBy:  createdBy,
		UpdatedBy:  createdBy,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.repo.Create(ctx, att); err != nil {
		return nil, errors.ErrInternalServer
	}

	return att, nil
}

func (s *attendanceService) GetAllAttendances(ctx context.Context) ([]attendance.Attendance, error) {
	return s.repo.GetAll(ctx)
}

func (s *attendanceService) GetAttendanceByID(ctx context.Context, id uint) (*attendance.Attendance, error) {
	att, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return att, nil
}

func (s *attendanceService) UpdateAttendance(ctx context.Context, id uint, employeeID, locationID uint, absentIn, absentOut *time.Time, updatedBy string) (*attendance.Attendance, error) {
	att, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	att.EmployeeID = employeeID
	att.LocationID = locationID
	att.AbsentIn = absentIn
	att.AbsentOut = absentOut
	att.UpdatedBy = updatedBy
	att.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, att); err != nil {
		return nil, errors.ErrInternalServer
	}

	return att, nil
}

func (s *attendanceService) DeleteAttendance(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.ErrInternalServer
	}
	return nil
}

func (s *attendanceService) GetAttendanceReport(ctx context.Context, startDate, endDate time.Time) ([]attendance.Attendance, error) {
	return s.repo.GetByDateRange(ctx, startDate, endDate)
}
