package rest

import (
	"net/http"

	"golang.org/x/net/context"
)

// itemDelete handles DELETE resquests on an item URL
func (r *request) itemDelete(ctx context.Context, route RouteMatch) (status int, headers http.Header, body interface{}) {
	lookup, e := route.Lookup()
	if e != nil {
		return e.Code, nil, e
	}
	l, err := route.Resource().Find(ctx, lookup, 1, 1)
	if err != nil {
		e = NewError(err)
		return e.Code, nil, e
	}
	if len(l.Items) == 0 {
		return ErrNotFound.Code, nil, ErrNotFound
	}
	original := l.Items[0]
	// If-Match / If-Unmodified-Since handling
	if err := r.checkIntegrityRequest(original); err != nil {
		return e.Code, nil, e
	}
	if err := route.Resource().Delete(ctx, original); err != nil {
		e = NewError(err)
		return e.Code, nil, e
	}
	return 204, nil, nil
}
