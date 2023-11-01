package middleware

import (
	"net/http"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/helpers"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/internal/usecases"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func Authorize(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace.Func()

		sessionId, ctx := helpers.SetGetAuthToContext(r)
		ctx = helpers.SetIsAdminCtx(ctx, false)

		if sessionId != "" {
			// todo improve
			sr := repos.NewSessionRepo(app.DB)
			mr := repos.NewMetadataRepo(app.DB)
			u := usecases.NewUserUc(mr, sr)
			err := u.AuthorizeUser(ctx)
			if err == nil {
				ctx = helpers.SetIsAdminCtx(ctx, true)
			}
		}
		nextHandler(w, r.WithContext(ctx))
	}
}
