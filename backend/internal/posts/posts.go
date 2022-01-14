package posts

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/users"
	"github.com/ganiszulfa/concise/backend/pkg/inspect"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/gosimple/slug"
)

var errMsgInputInvalid = "input is invalid"

func GetBySlug(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	id, ok := args["slug"].(int)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	result := app.DB.WithContext(ctx).First(&post, "slug = ?", id)
	return post, result.Error
}

func GetList(ctx context.Context, args map[string]interface{}) (posts []models.Post, err error) {
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
		Limit(limit).Offset(offset).Preload("Author").Find(&posts)

	return posts, result.Error
}

func Create(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Post{}, errors.New("login is required")
	}

	if !user.IsOwner {
		return models.Post{}, errors.New("forbidden")
	}

	title, ok := args["title"].(string)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	content, ok := args["content"].(string)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	post = models.Post{
		Title:    title,
		Content:  content,
		AuthorID: user.Id,
		Slug:     generateSlug(title),
	}

	result := app.DB.Create(&post)

	return post, result.Error
}

func Update(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Post{}, errors.New("login is required")
	}

	post, err = GetBySlug(ctx, args)
	if err != nil {
		return post, err
	}

	if user.Id != post.AuthorID {
		return models.Post{}, errors.New("forbidden")
	}

	title, ok := args["title"].(string)
	if ok {
		post.Title = title
		post.Slug = generateSlug(title)
	}

	content, ok := args["content"].(string)
	if ok {
		post.Content = content
	}

	return models.Post{}, nil
}

func generateSlug(s string) string {
	i := 10 * 1000
	r := fmt.Sprintf("_%d", rand.Intn(i*10)+i)
	x := slug.Make(s) + r
	inspect.Do(x)
	return x
}
