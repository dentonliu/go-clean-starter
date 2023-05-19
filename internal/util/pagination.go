package util

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/dentonliu/go-clean-starter/pkg"
)

const (
	DefaultPageSize int = 20
	MaxPageSize     int = 500
)

func GetPaginatedListFromRequest(c *gin.Context, count int) *pkg.PaginatedList {
	page := parseInt(c.Query("page"), 1)
	perPage := parseInt(c.Query("per_page"), DefaultPageSize)
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	return pkg.NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
