package api

import "math"

const defaultSize = 12
const maxSize = 60

// Pagination pagination data
type Pagination struct {
	PageNumber int32 `json:"pageNumber"`
	PageCount  int32 `json:"pageCount"`
	PageSize   int32 `json:"pageSize"`
	TotalCount int32 `json:"totalCount"`
}

// Limit limit
func (p *Pagination) Limit() int32 {
	return p.PageSize
}

// Offset offset
func (p *Pagination) Offset() int32 {
	offset := (p.PageNumber - 1) * p.PageSize
	if offset < 0 {
		return 0
	}
	return offset
}

// SetTotal set total and recalculated data
func (p *Pagination) SetTotal(total int32) {
	if total > 0 {
		pageCount := int32(math.Ceil(float64(total) / float64(p.PageSize)))
		pageNumber := p.PageNumber
		if p.PageNumber <= 0 {
			pageNumber = 1
		}

		p.PageCount = pageCount
		p.PageNumber = pageNumber
		p.TotalCount = total
	}
}

// InitPagination new pagination with pageNumber and size
func InitPagination(pageNumber int32, size int32) Pagination {
	if pageNumber <= 0 {
		pageNumber = 0
	}
	if size <= 0 {
		size = defaultSize
	}
	if size > maxSize {
		size = maxSize
	}
	return Pagination{
		PageNumber: pageNumber,
		PageSize:   size,
	}
}
