package posts

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/users"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/gosimple/slug"
)

var errMsgInputInvalid = "input is invalid"

func GetById(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	id, ok := args["id"].(int)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	isPublished, ok := args["isPublished"].(bool)
	if !ok {
		isPublished = true
	}

	if !isPublished {
		user, ok := users.GetUserFromCtx(ctx)
		if !ok || !user.IsOwner {
			isPublished = true
		}
	}

	return doGetById(ctx, id, &isPublished)
}

func doGetById(ctx context.Context, id int, isPublished *bool) (post models.Post, err error) {
	trace.Func()

	result := app.DB.WithContext(ctx)

	if isPublished != nil {
		result = result.
			Where("is_published", isPublished)
	}

	result = result.First(&post, "id = ?", id)

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

	isPublished, ok := args["isPublished"].(bool)
	if !ok {
		isPublished = true
	}

	if !isPublished {
		user, ok := users.GetUserFromCtx(ctx)
		if !ok || !user.IsOwner {
			isPublished = true
		}
	}

	result := app.DB.WithContext(ctx).
		Limit(limit).Offset(offset).Order(`"created_at" desc`).
		Where(`"is_published"`, isPublished).
		Preload("Author").Find(&posts)

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

	isPublished, ok := args["isPublished"].(bool)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	post = models.Post{
		Title:       title,
		Content:     content,
		AuthorID:    user.Id,
		IsPublished: isPublished,
		Slug:        slug.Make(title),
	}

	if isPublished {
		post.PublishedAt = time.Now()
	}

	result := app.DB.Create(&post)

	if result.Error != nil {
		post.Slug = generateSafePostSlug(title)
		result = app.DB.Create(&post)
	}

	return post, result.Error
}

func Update(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Post{}, errors.New("login is required")
	}

	id, ok := args["id"].(int)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	post, err = doGetById(ctx, id, nil)
	if err != nil {
		return post, err
	}

	if user.Id != post.AuthorID {
		return models.Post{}, errors.New("forbidden")
	}

	title, ok := args["title"].(string)
	if ok {
		post.Title = title
		post.Slug = slug.Make(title)
	}

	content, ok := args["content"].(string)
	if ok {
		post.Content = content
	}

	isPublished, ok := args["isPublished"].(bool)
	if ok {
		post.IsPublished = isPublished
	}

	if isPublished && post.PublishedAt.IsZero() {
		post.PublishedAt = time.Now()
	}

	result := app.DB.Model(&post).Updates(
		map[string]interface{}{
			"title":        post.Title,
			"slug":         post.Slug,
			"content":      post.Content,
			"is_published": post.IsPublished,
			"published_at": post.PublishedAt,
		},
	)

	if result.Error != nil {
		post.Slug = generateSafePostSlug(title)
		result = app.DB.Model(&post).Updates(
			map[string]interface{}{
				"title":        post.Title,
				"slug":         post.Slug,
				"content":      post.Content,
				"is_published": post.IsPublished,
				"published_at": post.PublishedAt,
			},
		)
	}

	return post, result.Error
}

func Delete(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Post{}, errors.New("login is required")
	}

	id, ok := args["id"].(int)
	if !ok {
		return models.Post{}, errors.New(errMsgInputInvalid)
	}

	post, err = doGetById(ctx, id, nil)
	if err != nil {
		return post, err
	}

	if user.Id != post.AuthorID {
		return models.Post{}, errors.New("forbidden")
	}

	result := app.DB.Delete(&post)

	return post, result.Error
}

func generateSafePostSlug(s string) string {
	i := 10 * 1000
	r := fmt.Sprintf("_%d", rand.Intn(i*10)+i)
	return slug.Make(s) + r
}
