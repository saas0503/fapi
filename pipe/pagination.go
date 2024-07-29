package pipe

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/saas0503/factory-api/exception"
)

type Token string

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
			exception.ThrowInvalidRequest(w, errors.New("invalid page parameter"))
			return
		}
		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			exception.ThrowInvalidRequest(w, errors.New("invalid limit parameter"))
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
