package usecases

import (
	"context"
	"errors"

	"github.com/ganiszulfa/concise/backend/internal/models/keys"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

type AuthorizationUcInterface interface {
	AuthorizeUser(ctx context.Context) (err error)
}

func NewAuthorizationUc(metadataRepo repos.MetadataRepoInterface) AuthorizationUcInterface {
	return &AuthorizationUc{metadataRepo: metadataRepo}
}

type AuthorizationUc struct {
	metadataRepo repos.MetadataRepoInterface
}

func (u AuthorizationUc) AuthorizeUser(ctx context.Context) (err error) {

	trace.Func()

	password, ok := ctx.Value("User-Password").(string)

	if !ok || password == "" {
		return errors.New("empty password")
	}

	mds, err := u.metadataRepo.GetAll(ctx)
	if err != nil {
		return
	}

	for _, v := range mds {
		if v.Key == keys.KEY_USER_PASSWORD {
			if password == v.Value {
				return nil
			} else {
				return errors.New("invalid password")
			}
		}
	}

	return errors.New("password key not found, invalid setup")
}
