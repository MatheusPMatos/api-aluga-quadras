package dto

import "math"

type Pagination struct {
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	Sort       string      `json:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 50
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {

	return (p.Offset / p.Limit) + 1
}
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func (p *Pagination) getTotalRows() int {
	if p.TotalRows == 0 {
		return 1
	}
	return int(p.TotalRows)
}

func (p *Pagination) GetTotalPages() int {
	totalPages := int(math.Ceil(float64(p.getTotalRows()) / float64(p.GetLimit())))
	return totalPages
}
