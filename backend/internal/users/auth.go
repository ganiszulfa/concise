package users

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/golang-jwt/jwt"
)

var userKey *bool

var ErrMsgTokenEmpty = "token is empty"
var SecretKey = []byte("secret")

type Claims struct {
	Id int `json:"id,omitempty"`
	jwt.StandardClaims
}

func VerifyTokenAndGetCtx(ctx context.Context, token string) (context.Context, error) {
	trace.Func()

	if token == "" {
		return ctx, errors.New(ErrMsgTokenEmpty)
	}

	s := strings.Fields(token)
	token = s[len(s)-1]

	claims, err := getClaimsFromToken(token)
	if err != nil {
		return ctx, err
	}

	user, err := getById(ctx, claims.Id)
	if err != nil {
		return ctx, err
	}

	newCtx := setUserAndGetNewCtx(ctx, user)

	return newCtx, nil
}

func getClaimsFromToken(tokenStr string) (*Claims, error) {
	trace.Func()

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, claims.Valid()
	} else {
		return nil, errors.New("invalid token")
	}
}

func getTokenForUser(user *models.User) (string, error) {
	trace.Func()

	sc := jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + int64(3600*24*14),
	}
	claims := Claims{
		Id:             int(user.Id),
		StandardClaims: sc,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString(SecretKey)
	return ss, err
}

func setUserAndGetNewCtx(ctx context.Context, user models.User) context.Context {
	trace.Func()

	return context.WithValue(ctx, userKey, user)
}

func GetUserFromCtx(ctx context.Context) (models.User, bool) {
	trace.Func()
	u, ok := ctx.Value(userKey).(models.User)
	return u, ok
}
