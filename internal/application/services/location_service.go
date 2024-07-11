package services

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/internal/domain/location"
	"github.com/andrMaulana/employee-management-api/pkg/errors"
)

type LocationService interface {
	CreateLocation(ctx context.Context, name, createdBy string) (*location.Location, error)
	GetAllLocations(ctx context.Context) ([]location.Location, error)
	GetLocationByID(ctx context.Context, id uint) (*location.Location, error)
	UpdateLocation(ctx context.Context, id uint, name, updatedBy string) (*location.Location, error)
	DeleteLocation(ctx context.Context, id uint) error
}

type locationService struct {
	repo location.Repository
}

func NewLocationService(repo location.Repository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) CreateLocation(ctx context.Context, name, createdBy string) (*location.Location, error) {
	loc := &location.Location{
		LocationName: name,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, loc); err != nil {
		return nil, errors.ErrInternalServer
	}

	return loc, nil
}

func (s *locationService) GetAllLocations(ctx context.Context) ([]location.Location, error) {
	return s.repo.GetAll(ctx)
}

func (s *locationService) GetLocationByID(ctx context.Context, id uint) (*location.Location, error) {
	loc, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return loc, nil
}

func (s *locationService) UpdateLocation(ctx context.Context, id uint, name, updatedBy string) (*location.Location, error) {
	loc, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	loc.LocationName = name
	loc.UpdatedBy = updatedBy
	loc.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, loc); err != nil {
		return nil, errors.ErrInternalServer
	}

	return loc, nil
}

func (s *locationService) DeleteLocation(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.ErrInternalServer
	}
	return nil
}
