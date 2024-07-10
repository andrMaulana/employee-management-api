// pkg/pagination/pagination.go
package pagination

type Paginator struct {
	Page  int `json:"page" query:"page"`
	Limit int `json:"limit" query:"limit"`
}

func (p *Paginator) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Paginator) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Paginator) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

type PagedResult[T any] struct {
	Data       []T   `json:"data"`
	TotalItems int64 `json:"total_items"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}

func NewPagedResult[T any](data []T, totalItems int64, page, limit int) *PagedResult[T] {
	totalPages := int(totalItems) / limit
	if int(totalItems)%limit > 0 {
		totalPages++
	}

	return &PagedResult[T]{
		Data:       data,
		TotalItems: totalItems,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}
