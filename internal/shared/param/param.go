package param

import "time"

type QueryParams struct {
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
	Page    int      `form:"page" param:"page" db:"page"`
	Limit   int      `form:"limit" param:"limit" db:"limit"`
	SortBy  []string `form:"sortBy" param:"sort_by" db:"sort_by"`
	GroupBy []string `form:"sortBy" param:"sort_by" db:"sort_by"`
}
