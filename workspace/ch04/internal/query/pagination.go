package query

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page    int
	PerPage int
	Offset  int
}

func ParsePagination(r *http.Request) Pagination {
	q := r.URL.Query()
	page := 1
	if p := q.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	perPage := 20
	if p := q.Get("per_page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			if v > 100 {
				v = 100
			}
			perPage = v
		}
	}
	return Pagination{
		Page:    page,
		PerPage: perPage,
		Offset:  (page - 1) * perPage,
	}
}
