package pipe

import (
	"context"
	"net/http"
	"strconv"
)

const PaginationToken Token = "pagination"

type Paginate struct {
	Limit  int
	Offset int
}

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")

		if page == "" {
			page = "1"
		}
		if limit == "" {
			limit = "10"
		}

		intPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		offset := (intPage - 1) * intLimit

		paginate := &Paginate{
			Limit:  intLimit,
			Offset: offset,
		}

		ctx := context.WithValue(r.Context(), PaginationToken, paginate)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
