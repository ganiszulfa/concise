package users

import (
	"context"
	"errors"
	"time"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/pwd"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/ganiszulfa/concise/backend/pkg/validator"
)

var errMsgInputInvalid = "input is invalid"

func getById(ctx context.Context, id int) (user models.User, err error) {
	trace.Func()

	result := app.DB.WithContext(ctx).First(&user, "id = ?", id)
	return user, result.Error
}

func GetByEmail(ctx context.Context, args map[string]interface{}) (user models.User, err error) {
	trace.Func()

	email, ok := args["email"].(string)
	if !ok {
		return models.User{}, errors.New(errMsgInputInvalid)
	}

	result := app.DB.WithContext(ctx).First(&user, "email = ?", email)

	return user, result.Error
}

func GetList(ctx context.Context, args map[string]interface{}) (users []models.User, err error) {
	trace.Func()

	page, ok := args["page"].(int)
	if !ok || page == 0 {
		page = 1
	}

	limit, ok := args["limit"].(int)
	if !ok || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	result := app.DB.WithContext(ctx).
		Limit(limit).Offset(offset).Find(&users)

	return users, result.Error
}

func Create(ctx context.Context, args map[string]interface{}) (user models.User, err error) {
	trace.Func()

	email, ok := args["email"].(string)
	if !ok || validator.IsNotValidEmail(email) {
		return models.User{}, errors.New(errMsgInputInvalid)
	}

	password, ok := args["password"].(string)
	if !ok || len(password) < 4 {
		return models.User{}, errors.New(errMsgInputInvalid)
	}

	firstName, _ := args["firstName"].(string)
	lastName, _ := args["lastName"].(string)
	isOwner, ok := args["isOwner"].(bool)

	if !ok {
		isOwner = false
	}

	if isOwner {
		ownerPassword, ok := args["ownerPassword"].(string)
		if !ok {
			return models.User{}, errors.New("owner password is required")
		}

		// to reduce bruteforce
		time.Sleep(time.Duration(100) * time.Millisecond)

		if ownerPassword != app.OwnerPassword {
			return models.User{}, errors.New("invalid owner password")
		}
	}

	hashedPassword, err := pwd.Hash(password)
	if err != nil {
		panic(err)
	}

	user = models.User{
		Email:     email,
		Password:  hashedPassword,
		FirstName: firstName,
		LastName:  lastName,
		IsOwner:   isOwner,
	}

	result := app.DB.Create(&user)

	return user, result.Error
}

func Update(ctx context.Context, args map[string]interface{}) (user models.User, err error) {
	trace.Func()

	user, ok := GetUserFromCtx(ctx)
	if !ok {
		return models.User{}, errors.New("login is required")
	}

	email, ok := args["email"].(string)
	if ok && validator.IsNotValidEmail(email) {
		user.Email = email
	}

	password, ok := args["password"].(string)
	if ok {
		if len(password) < 4 {
			return models.User{}, errors.New("password is too short")
		}
		hashedPassword, err := pwd.Hash(password)
		if err != nil {
			panic(err)
		}
		user.Password = hashedPassword
	}

	firstName, ok := args["firstName"].(string)
	if ok {
		user.FirstName = firstName
	}

	lastName, ok := args["lastName"].(string)
	if ok {
		user.LastName = lastName
	}

	return models.User{}, nil
}

func Login(ctx context.Context, args map[string]interface{}) (user models.User, err error) {
	trace.Func()

	user, err = GetByEmail(ctx, args)
	if err != nil {
		return models.User{}, errors.New("invalid email/password")
	}

	password, ok := args["password"].(string)
	if !ok {
		return models.User{}, errors.New(errMsgInputInvalid)
	}

	if !pwd.CheckHash(password, user.Password) {
		return models.User{}, errors.New("invalid email/password")
	}

	user.Token, err = getTokenForUser(&user)
	return user, err
}
