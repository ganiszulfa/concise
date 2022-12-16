package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func (c PostCtr) GetBySlugFromGQL(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	slug, ok := args[ArgsPostsSlug].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	isPublished, ok := args[ArgsPostsIsPublished].(bool)
	if !ok {
		isPublished = true
	}

	return c.postUc.GetBySlug(ctx, slug, &isPublished)
}

func (c PostCtr) GetListFromGQL(ctx context.Context, args map[string]interface{}) (posts []models.Post, err error) {
	trace.Func()

	page, ok := args[ArgsPage].(int)
	if !ok || page == 0 {
		page = 1
	}

	limit, ok := args[ArgsLimit].(int)
	if !ok || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	isPage, ok := args[ArgsPostsIsPage].(bool)
	if !ok {
		isPage = false
	}

	isPublished, ok := args[ArgsPostsIsPublished].(bool)
	if !ok {
		isPublished = true
	}

	return c.postUc.GetList(ctx, limit, offset, &isPage, &isPublished)
}
