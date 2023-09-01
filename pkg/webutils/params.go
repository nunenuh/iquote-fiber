package webutils

import "github.com/gofiber/fiber/v2"

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
