package middleware

import (
	"context"
	"net/http"

	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func Authorize(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace.Func()
		ctx := context.WithValue(r.Context(), "User-Password", r.Header.Get("User-Password"))
		nextHandler(w, r.WithContext(ctx))
	}
}
