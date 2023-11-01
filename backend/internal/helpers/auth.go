package helpers

import (
	"context"
	"errors"
	"net/http"
)

var SESSION_ID_KEY = "SessionId"
var SESSION_ID_CONTEXT_KEY struct{}
var IS_ADMIN_CONTEXT_KEY struct{}

func IsAdmin(ctx context.Context) (bool, error) {
	isAdmin, ok := ctx.Value(IS_ADMIN_CONTEXT_KEY).(bool)
	if ok && isAdmin {
		return true, nil
	}

	return false, errors.New("not authorized")
}

func GetSessionId(ctx context.Context) (string, error) {
	sessionId, ok := ctx.Value(SESSION_ID_CONTEXT_KEY).(string)

	if !ok || sessionId == "" {
		return "", errors.New("empty session id")
	}
	return sessionId, nil
}

func SetGetAuthToContext(r *http.Request) (string, context.Context) {
	sessionId := r.Header.Get(SESSION_ID_KEY)
	ctx := context.WithValue(r.Context(), SESSION_ID_CONTEXT_KEY, sessionId)
	return sessionId, ctx
}

func SetIsAdminCtx(ctx context.Context, val bool) context.Context {
	return context.WithValue(ctx, IS_ADMIN_CONTEXT_KEY, val)
}
