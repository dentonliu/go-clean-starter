package pkg

const (
	DefaultPageSize int = 20
)

type PaginatedList struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

func NewPaginatedList(page, perPage, total int) *PaginatedList {
	if perPage < 1 {
		perPage = DefaultPageSize
	}
	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &PaginatedList{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

func (p *PaginatedList) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *PaginatedList) Limit() int {
	return p.PerPage
}
