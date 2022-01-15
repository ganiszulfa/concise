package pages

import (
	"context"
	"errors"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/users"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/gosimple/slug"
)

var errMsgInputInvalid = "input is invalid"

func GetBySlug(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	slug, ok := args["slug"].(string)
	if !ok {
		return models.Page{}, errors.New(errMsgInputInvalid)
	}

	result := app.DB.WithContext(ctx).First(&page, "slug = ?", slug)
	return page, result.Error
}

func GetList(ctx context.Context, args map[string]interface{}) (pages []models.Page, err error) {
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
		Limit(limit).Offset(offset).Order(`"order", "title"`).
		Preload("Author").Find(&pages)

	return pages, result.Error
}

func Create(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Page{}, errors.New("login is required")
	}

	if !user.IsOwner {
		return models.Page{}, errors.New("forbidden")
	}

	title, ok := args["title"].(string)
	if !ok {
		return models.Page{}, errors.New(errMsgInputInvalid)
	}

	content, ok := args["content"].(string)
	if !ok {
		return models.Page{}, errors.New(errMsgInputInvalid)
	}

	order, ok := args["order"].(int)
	if !ok {
		return models.Page{}, errors.New(errMsgInputInvalid)
	}

	page = models.Page{
		Title:    title,
		Content:  content,
		Order:    order,
		AuthorID: user.Id,
		Slug:     slug.Make(title),
	}

	result := app.DB.Create(&page)

	return page, result.Error
}

func Update(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Page{}, errors.New("login is required")
	}

	page, err = GetBySlug(ctx, args)
	if err != nil {
		return page, err
	}

	if user.Id != page.AuthorID {
		return models.Page{}, errors.New("forbidden")
	}

	title, ok := args["title"].(string)
	if ok {
		page.Title = title
		page.Slug = slug.Make(title)
	}

	content, ok := args["content"].(string)
	if ok {
		page.Content = content
	}

	result := app.DB.Model(&page).Updates(page)

	return page, result.Error
}

func Delete(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	user, ok := users.GetUserFromCtx(ctx)
	if !ok {
		return models.Page{}, errors.New("login is required")
	}

	page, err = GetBySlug(ctx, args)
	if err != nil {
		return page, err
	}

	if user.Id != page.AuthorID {
		return models.Page{}, errors.New("forbidden")
	}

	result := app.DB.Delete(&page)

	return page, result.Error
}
