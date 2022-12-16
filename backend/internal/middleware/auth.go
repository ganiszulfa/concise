package middleware

import (
	"net/http"

	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func Authorize(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace.Func()

		nextHandler(w, r.WithContext(r.Context()))
	}
}
