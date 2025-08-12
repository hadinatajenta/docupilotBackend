package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const DefaultPerPage = 20
const MaxPerPage = 100

type Params struct {
	Page    int
	PerPage int
	Sort    string // e.g. "created_at"
	Order   string // "asc"/"desc"
	Offset  int
	Limit   int
}

type Meta struct {
	TotalItems  int  `json:"total_items,omitempty"`
	CurrentPage int  `json:"current_page,omitempty"`
	PerPage     int  `json:"per_page,omitempty"`
	TotalPages  int  `json:"total_pages,omitempty"`
	HasNext     bool `json:"has_next,omitempty"`
	HasPrev     bool `json:"has_prev,omitempty"`
}

func BuildMeta(total, page, per int) Meta {
	tp := (total + per - 1) / per
	return Meta{
		TotalItems:  total,
		CurrentPage: page,
		PerPage:     per,
		TotalPages:  tp,
		HasNext:     page < tp,
		HasPrev:     page > 1,
	}
}

func Parse(c *gin.Context) Params {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", strconv.Itoa(DefaultPerPage)))
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	if page < 1 {
		page = 1
	}
	if per < 1 {
		per = DefaultPerPage
	}
	if per > MaxPerPage {
		per = MaxPerPage
	}
	// sanitize order
	if order != "asc" && order != "desc" {
		order = "desc"
	}

	return Params{
		Page:    page,
		PerPage: per,
		Sort:    sort,
		Order:   order,
		Offset:  (page - 1) * per,
		Limit:   per,
	}
}
