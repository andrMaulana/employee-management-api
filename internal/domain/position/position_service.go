package position

import (
	"context"
	"time"

	"github.com/andrMaulana/employee-management-api/pkg/pagination"
)

type Service interface {
	Create(ctx context.Context, position *Position) error
	GetAll(ctx context.Context, params *GetAllParams) (*pagination.PagedResult[Position], error)
	GetByID(ctx context.Context, id uint) (*Position, error)
	Update(ctx context.Context, position *Position) error
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

func (s *service) Create(ctx context.Context, position *Position) error {
	position.CreatedAt = time.Now()
	position.UpdatedAt = time.Now()
	return s.repo.Create(ctx, position)
}

func (s *service) GetAll(ctx context.Context, params *GetAllParams) (*pagination.PagedResult[Position], error) {
	repoParams := &FindAllParams{
		Paginator: params.Paginator,
		Search:    params.Search,
		SortBy:    params.SortBy,
		SortOrder: params.SortOrder,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}

	positions, total, err := s.repo.FindAll(ctx, repoParams)
	if err != nil {
		return nil, err
	}

	return &pagination.PagedResult[Position]{
		Data:       positions,
		TotalItems: total,
		Page:       params.GetPage(),
		Limit:      params.GetLimit(),
	}, nil
}

func (s *service) GetByID(ctx context.Context, id uint) (*Position, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, position *Position) error {
	position.UpdatedAt = time.Now()
	return s.repo.Update(ctx, position)
}

func (s *service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
