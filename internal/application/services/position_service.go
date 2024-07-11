package services

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/position"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
)

type PositionService interface {
	CreatePosition(ctx context.Context, name string, departmentID uint, createdBy string) (*position.Position, error)
	GetAllPositions(ctx context.Context) ([]position.Position, error)
	GetPositionByID(ctx context.Context, id uint) (*position.Position, error)
	UpdatePosition(ctx context.Context, id uint, name string, departmentID uint, updatedBy string) (*position.Position, error)
	DeletePosition(ctx context.Context, id uint) error
}

type positionService struct {
	repo position.Repository
}

func NewPositionService(repo position.Repository) PositionService {
	return &positionService{repo: repo}
}

func (s *positionService) CreatePosition(ctx context.Context, name string, departmentID uint, createdBy string) (*position.Position, error) {
	pos := &position.Position{
		PositionName: name,
		DepartmentID: departmentID,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, pos); err != nil {
		return nil, errors.ErrInternalServer
	}

	return pos, nil
}

func (s *positionService) GetAllPositions(ctx context.Context) ([]position.Position, error) {
	return s.repo.GetAll(ctx)
}

func (s *positionService) GetPositionByID(ctx context.Context, id uint) (*position.Position, error) {
	pos, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return pos, nil
}

func (s *positionService) UpdatePosition(ctx context.Context, id uint, name string, departmentID uint, updatedBy string) (*position.Position, error) {
	pos, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	pos.PositionName = name
	pos.DepartmentID = departmentID
	pos.UpdatedBy = updatedBy
	pos.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, pos); err != nil {
		return nil, errors.ErrInternalServer
	}

	return pos, nil
}

func (s *positionService) DeletePosition(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.ErrInternalServer
	}
	return nil
}
