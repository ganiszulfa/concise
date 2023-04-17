package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/responses"

	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func (c UserCtr) UserLoginFromGQL(ctx context.Context, args map[string]interface{}) (loginResponse responses.Login, err error) {
	trace.Func()

	id, ok := args[ArgsUserId].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	password, ok := args[ArgsUserPassword].(string)
	if !ok {
		err = errInputInvalid
		return
	}
	return c.userUc.Login(ctx, id, password)
}
