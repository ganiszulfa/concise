package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/responses"

	"github.com/ganiszulfa/concise/backend/internal/usecases"
)

const (
	ObjectNameUser      = "User"
	ObjectNameUserLogin = "UserLogin"

	UserQueryLogin = "Login"
)

const (
	ArgsUserId       = "id"
	ArgsUserPassword = "password"
	ArgsUserSession  = "sessionId"
)

type UserCtrInterface interface {
	UserLoginFromGQL(ctx context.Context, args map[string]interface{}) (loginResponse responses.Login, err error)
}

type UserCtr struct {
	userUc usecases.UserUcInterface
}

func NewUserCtr(userUc usecases.UserUcInterface) UserCtrInterface {
	return &UserCtr{userUc: userUc}
}
