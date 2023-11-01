package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/ganiszulfa/concise/backend/internal/helpers"
	"github.com/ganiszulfa/concise/backend/internal/models/keys"

	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/internal/responses"
	"github.com/ganiszulfa/concise/backend/pkg/pwd"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

var SESSION_COOKIE_NAME = "session"

type UserUcInterface interface {
	Login(ctx context.Context, id, password string) (loginResponse responses.Login, err error)
	AuthorizeUser(ctx context.Context) (err error)
}

type UserUc struct {
	metadataRepo repos.MetadataRepoInterface
	sessionRepo  repos.SessionRepoInterface
}

func NewUserUc(
	metadataRepo repos.MetadataRepoInterface,
	sessionRepo repos.SessionRepoInterface,
) UserUcInterface {

	return &UserUc{
		metadataRepo: metadataRepo,
		sessionRepo:  sessionRepo,
	}
}

func (u UserUc) Login(ctx context.Context, id, password string) (loginResponse responses.Login, err error) {

	trace.Func()

	mdHashedPassword, err := u.metadataRepo.GetByKey(ctx, keys.KEY_USER_PASSWORD)
	if err != nil {
		return
	}

	if !pwd.CheckHash(password, mdHashedPassword.Value) {
		err = errors.New("invalid password")
		return
	}

	sessionId, err := u.sessionRepo.Create(ctx)
	if err != nil {
		return
	}

	loginResponse.SessionId = sessionId
	return
}

func (u UserUc) AuthorizeUser(ctx context.Context) (err error) {

	trace.Func()

	sessionId, err := helpers.GetSessionId(ctx)
	if err != nil {
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
