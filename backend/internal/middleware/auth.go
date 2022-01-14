package middleware

import (
	"net/http"

	"github.com/ganiszulfa/concise/internal/users"
	"github.com/ganiszulfa/concise/pkg/trace"
)

func Authorize(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace.Func()

		token := r.Header.Get("Authorization")
		ctx := r.Context()

		newCtx, err := users.VerifyTokenAndGetCtx(ctx, token)

		if err != nil && err.Error() != users.ErrMsgTokenEmpty {
			http.Error(w, err.Error(), http.StatusForbidden)
		}

		nextHandler(w, r.WithContext(newCtx))
	}
}
