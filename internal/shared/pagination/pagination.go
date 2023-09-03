package pagination

import param "github.com/nunenuh/iquote-fiber/internal/shared/param"

type Pagination struct {
	CurrentPage     int      `json:"currentPage"`
	CurrentElements int      `json:"currentElements"`
	TotalPages      int      `json:"totalPages"`
	TotalElements   int      `json:"totalElements"`
	SortBy          []string `json:"sortBy"`
	CursorStart     *string  `json:"cursorStart,omitempty"`
	CursorEnd       *string  `json:"cursorEnd,omitempty"`
}

func NewPagination(param param.Param, totalElements int) *Pagination {
	p := &Pagination{
		CurrentPage:     int(param.Page),
		CurrentElements: totalElements, // Assuming you want the length of authors as the current number of elements
		SortBy:          param.SortBy,
		TotalElements:   totalElements,
	}
	p.ProcessPagination(param.Limit)
	return p
}

func (p *Pagination) ProcessPagination(limit int) {
	if p.SortBy == nil {
		p.SortBy = []string{}
	}
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalPage := p.TotalElements / limit
	if p.TotalElements%limit > 0 || p.TotalElements == 0 {
		totalPage++
	}

	p.TotalPages = 1
	if totalPage > 1 {
		p.TotalPages = totalPage
	}
}
