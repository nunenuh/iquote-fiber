package webutils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type QueryParameters struct {
	Limit             int64
	Offset            int64
	OrderBy           string
	OrderDirection    string
	Fields            []string
	Filters           map[string]interface{}
	Search            string
	IncludeDeleted    bool
	LastModifiedSince time.Time
	GroupBy           []string
}

type Param struct {
	Page   int      `form:"page" param:"page" db:"page"`
	Limit  int      `form:"limit" param:"limit" db:"limit"`
	SortBy []string `form:"sortBy" param:"sort_by" db:"sort_by"`
}

type PaginationParam struct {
	GroupBy []string `param:"-" db:"-"`
	SortBy  []string `param:"sort_by" db:"sort_by"`
	Limit   int64    `form:"limit" param:"limit" db:"limit"`
	Page    int64    `form:"page" param:"page" db:"page"`
}

func ParsePaginationParams(c *fiber.Ctx) (*PaginationParam, error) {
	params := &PaginationParam{}
	if err := c.QueryParser(params); err != nil {
		return nil, err
	}
	return params, nil
}
