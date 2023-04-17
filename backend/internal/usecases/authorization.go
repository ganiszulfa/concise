package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

type AuthorizationUcInterface interface {
	AuthorizeUser(ctx context.Context) (err error)
}

func NewAuthorizationUc(
	sessionRepo repos.SessionRepoInterface,
) AuthorizationUcInterface {

	return &AuthorizationUc{
		sessionRepo: sessionRepo,
	}
}

type AuthorizationUc struct {
	sessionRepo repos.SessionRepoInterface
}

func (u AuthorizationUc) AuthorizeUser(ctx context.Context) (err error) {

	trace.Func()

	sessionId, ok := ctx.Value("SessionId").(string)

	if !ok || sessionId == "" {
		return errors.New("empty session id")
	}

	session, err := u.sessionRepo.GetById(ctx, sessionId)
	if err != nil {
		return
	}
	if session.ExpiredAt.Before(time.Now()) {
		return errors.New("session expired")
	}

	return nil
}
