package location

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/pkg/pagination"
)

type Service interface {
	Create(ctx context.Context, location *Location) error
	GetAll(ctx context.Context, params *GetAllParams) (*pagination.PagedResult[Location], error)
	GetByID(ctx context.Context, id uint) (*Location, error)
	Update(ctx context.Context, location *Location) error
	Delete(ctx context.Context, id uint) error
}
type service struct {
	repo Repository
}

type GetAllParams struct {
	pagination.Paginator
	Search    string
	SortBy    string
	SortOrder string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, location *Location) error {
	location.CreatedAt = time.Now()
	location.UpdatedAt = time.Now()
	return s.repo.Create(ctx, location)
}

func (s *service) GetAll(ctx context.Context, params *GetAllParams) (*pagination.PagedResult[Location], error) {
	repoParams := &FindAllParams{
		Paginator: params.Paginator,
		Search:    params.Search,
		SortBy:    params.SortBy,
		SortOrder: params.SortOrder,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}

	location, total, err := s.repo.FindAll(ctx, repoParams)
	if err != nil {
		return nil, err
	}

	return &pagination.PagedResult[Location]{
		Data:       location,
		TotalItems: total,
		Page:       params.GetPage(),
		Limit:      params.GetLimit(),
	}, nil
}

func (s *service) GetByID(ctx context.Context, id uint) (*Location, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, location *Location) error {
	location.UpdatedAt = time.Now()
	return s.repo.Update(ctx, location)
}

func (s *service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
