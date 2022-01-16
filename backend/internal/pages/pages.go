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

func GetById(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	id, ok := args["id"].(int)
	if !ok {
		return models.Page{}, errors.New(errMsgInputInvalid)
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

func doGetById(ctx context.Context, id int, isPublished *bool) (page models.Page, err error) {
	trace.Func()

	result := app.DB.WithContext(ctx)

	if isPublished != nil {
		result = result.
			Where("is_published", isPublished)
	}

	result = result.First(&page, "id = ?", id)

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
		Limit(limit).Offset(offset).Order(`"order", "title"`).
		Where(`"is_published"`, isPublished).
		Preload("Author").Find(&pages)

	return pages, result.Error
}

func Create(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	user, err := users.CheckIfOwner(ctx)
	if err != nil {
		return page, err
	}

	title, ok := args["title"].(string)
	if !ok {
		return page, errors.New(errMsgInputInvalid)
	}

	content, ok := args["content"].(string)
	if !ok {
		return page, errors.New(errMsgInputInvalid)
	}

	order, ok := args["order"].(int)
	if !ok {
		return page, errors.New(errMsgInputInvalid)
	}

	isPublished, ok := args["isPublished"].(bool)
	if !ok {
		return page, errors.New(errMsgInputInvalid)
	}

	page = models.Page{
		Title:       title,
		Content:     content,
		Order:       order,
		AuthorID:    user.Id,
		IsPublished: isPublished,
		Slug:        slug.Make(title),
	}

	result := app.DB.Create(&page)

	return page, result.Error
}

func Update(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	_, err = users.CheckIfOwner(ctx)
	if err != nil {
		return page, err
	}

	id, ok := args["id"].(int)
	if !ok {
		return page, errors.New(errMsgInputInvalid)
	}

	page, err = doGetById(ctx, id, nil)
	if err != nil {
		return page, err
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

	order, ok := args["order"].(int)
	if ok {
		page.Order = order
	}

	isPublished, ok := args["isPublished"].(bool)
	if ok {
		page.IsPublished = isPublished
	}

	result := app.DB.Model(&page).Updates(
		map[string]interface{}{
			"order":        page.Order,
			"title":        page.Title,
			"slug":         page.Slug,
			"content":      page.Content,
			"is_published": page.IsPublished,
		},
	)

	return page, result.Error
}

func Delete(ctx context.Context, args map[string]interface{}) (page models.Page, err error) {
	trace.Func()

	_, err = users.CheckIfOwner(ctx)
	if err != nil {
		return page, err
	}

	id, ok := args["id"].(int)
	if !ok {
		return models.Page{}, errors.New(errMsgInputInvalid)
	}

	page, err = doGetById(ctx, id, nil)
	if err != nil {
		return page, err
	}

	result := app.DB.Delete(&page)

	return page, result.Error
}
