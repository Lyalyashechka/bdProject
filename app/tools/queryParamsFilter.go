package tools

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

const (
	NameLimitParam = "limit"
	NameDescParam  = "desc"
	NameSinceParam = "since"
)

const (
	LimitParamDefault = 100
	SinceParamDefault = ""
	SortParamDefault  = "asc"
	SortParamTrue     = "desc"
)

type Filter struct {
	Limit int
	Sort  string
	Since string
}

func ParseQueryFilter(ctx echo.Context) Filter {
	var result Filter
	queryParam := ctx.QueryParams()

	limit := queryParam.Get(NameLimitParam)
	if limit != "" {
		limitInt, err := strconv.ParseInt(limit, 10, 32)
		if err != nil {
			result.Limit = 100
		} else {
			result.Limit = int(limitInt)
		}
	} else {
		result.Limit = LimitParamDefault
	}

	sort := queryParam.Get(NameDescParam)
	if sort == "true" {
		result.Sort = SortParamTrue
	} else {
		result.Sort = SortParamDefault
	}

	since := queryParam.Get(NameSinceParam)
	result.Since = since

	return result
}
