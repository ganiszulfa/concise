package usecases

import (
	"context"
	"errors"

	"github.com/ganiszulfa/concise/backend/internal/models/keys"

	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/internal/responses"
	"github.com/ganiszulfa/concise/backend/pkg/pwd"
)

var SESSION_COOKIE_NAME = "session"

type UserUcInterface interface {
	Login(ctx context.Context, id, password string) (loginResponse responses.Login, err error)
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
