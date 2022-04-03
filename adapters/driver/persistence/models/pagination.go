package models

import "fmt"

type Pageable interface {
	PageNumber() int
	Offset() int
	PageSize() int
}

type Pagination struct {
	Page      int    `json:"page" uri:"page" form:"page"`
	Limit     int    `json:"limit" uri:"limit" form:"limit"`
	Field     string `json:"field" uri:"field" form:"field"`
	Direction string `json:"direction" uri:"direction" form:"direction"` // TODO, validation, default
}

func NewPaginationWithDefaults() Pagination {
	return Pagination{Page: 0, Limit: 100}
}

func (p Pagination) PageNumber() int {
	return p.Page
}

func (p Pagination) Offset() int {
	return p.Page * p.Limit
}

func (p Pagination) PageSize() int {
	return p.Limit
}

func (p Pagination) IsOrderDefined() bool {
	return p.Field != "" && p.Direction != ""
}

func (p Pagination) Order() string {
	if !p.IsOrderDefined() {
		return ""
	}

	return fmt.Sprintf("%s %s", p.Field, p.Direction)
}
