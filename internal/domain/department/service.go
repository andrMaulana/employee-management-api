package department

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/pkg/pagination"
)

type Service interface {
	Create(ctx context.Context, department *Department) error
	GetAll(ctx context.Context, params *GetAllParams) (*pagination.PagedResult[Department], error)
	GetByID(ctx context.Context, id uint) (*Department, error)
	Update(ctx context.Context, department *Department) error
	Delete(ctx context.Context, id uint) error
	BatchCreate(ctx context.Context, departments []Department) error
	BatchUpdate(ctx context.Context, departments []Department) error
	BatchDelete(ctx context.Context, ids []uint) error
}

type GetAllParams struct {
	pagination.Paginator
	Search    string
	SortBy    string
	SortOrder string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, department *Department) error {
	department.CreatedAt = time.Now()
	department.UpdatedAt = time.Now()
	return s.repo.Create(ctx, department)
}

func (s *service) GetAll(ctx context.Context, params *GetAllParams) (*pagination.PagedResult[Department], error) {
	repoParams := &FindAllParams{
		Paginator: params.Paginator,
		Search:    params.Search,
		SortBy:    params.SortBy,
		SortOrder: params.SortOrder,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}

	departments, total, err := s.repo.FindAll(ctx, repoParams)
	if err != nil {
		return nil, err
	}

	return &pagination.PagedResult[Department]{
		Data:       departments,
		TotalItems: total,
		Page:       params.GetPage(),
		Limit:      params.GetLimit(),
	}, nil
}

func (s *service) GetByID(ctx context.Context, id uint) (*Department, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, department *Department) error {
	department.UpdatedAt = time.Now()
	return s.repo.Update(ctx, department)
}

func (s *service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) BatchCreate(ctx context.Context, departments []Department) error {
	now := time.Now()
	for i := range departments {
		departments[i].CreatedAt = now
		departments[i].UpdatedAt = now
	}
	return s.repo.BatchCreate(ctx, departments)
}

func (s *service) BatchUpdate(ctx context.Context, departments []Department) error {
	now := time.Now()
	for i := range departments {
		departments[i].UpdatedAt = now
	}
	return s.repo.BatchUpdate(ctx, departments)
}

func (s *service) BatchDelete(ctx context.Context, ids []uint) error {
	return s.repo.BatchDelete(ctx, ids)
}
