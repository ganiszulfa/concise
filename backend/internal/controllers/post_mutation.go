package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func (c PostCtr) CreateFromGQL(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	title, ok := args[ArgsPostsTitle].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	content, ok := args[ArgsPostsContent].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	isPublished, ok := args[ArgsPostsIsPublished].(bool)
	if !ok {
		isPublished = true
	}

	isPage, ok := args[ArgsPostsIsPage].(bool)
	if !ok {
		isPublished = true
	}

	return c.postUc.Create(ctx, title, content, isPage, isPublished)
}

func (c PostCtr) UpdateFromGQL(ctx context.Context, args map[string]interface{}) (post models.Post, err error) {
	trace.Func()

	id, ok := args[ArgsPostsId].(int)
	if !ok {
		err = errInputInvalid
		return
	}

	slug, ok := args[ArgsPostsSlug].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	title, ok := args[ArgsPostsTitle].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	content, ok := args[ArgsPostsContent].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	isPublished, ok := args[ArgsPostsIsPublished].(bool)
	if !ok {
		isPublished = true
	}

	return c.postUc.Update(ctx, id, slug, title, content, isPublished)
}

func (c PostCtr) DeleteFromGQL(ctx context.Context, args map[string]interface{}) (err error) {
	trace.Func()

	slug, ok := args[ArgsPostsSlug].(string)
	if !ok {
		err = errInputInvalid
		return
	}

	return c.postUc.Delete(ctx, slug)
}
