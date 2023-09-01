package webutils

type Pagination struct {
	CurrentPage     int64    `json:"currentPage"`
	CurrentElements int64    `json:"currentElements"`
	TotalPages      int64    `json:"totalPages"`
	TotalElements   int64    `json:"totalElements"`
	SortBy          []string `json:"sortBy"`
	CursorStart     *string  `json:"cursorStart,omitempty"`
	CursorEnd       *string  `json:"cursorEnd,omitempty"`
}

func NewPagination(params *PaginationParam, totalElements int64) *Pagination {
	p := &Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(totalElements), // Assuming you want the length of authors as the current number of elements
		SortBy:          params.SortBy,
		TotalElements:   totalElements,
	}
	p.ProcessPagination(params.Limit)
	return p
}

func (p *Pagination) ProcessPagination(limit int64) {
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
