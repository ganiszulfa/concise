package usecases

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

type PostUcInterface interface {
	GetBySlug(ctx context.Context, slug string, isPublished *bool) (post models.Post, err error)
	GetList(ctx context.Context, limit, offset int, isPage, isPublished *bool) (posts []models.Post, err error)
	Create(ctx context.Context, title, content string, isPage, isPublished bool) (post models.Post, err error)
	Update(ctx context.Context, id int, slug, title, content string, isPublished bool) (post models.Post, err error)
	Delete(ctx context.Context, slug string) (err error)
}

type PostUc struct {
	postRepo        repos.PostRepoInterface
	AuthorizationUc AuthorizationUcInterface
}

func NewPostUc(postRepo repos.PostRepoInterface,
	authoAuthorizationUc AuthorizationUcInterface) PostUcInterface {
	return &PostUc{postRepo: postRepo, AuthorizationUc: authoAuthorizationUc}
}

func (u PostUc) GetBySlug(ctx context.Context, slug string, isPublished *bool) (post models.Post, err error) {

	trace.Func()

	return u.postRepo.GetBySlug(ctx, slug, isPublished)
}

func (u PostUc) GetList(ctx context.Context, limit, offset int, isPage, isPublished *bool) (posts []models.Post, err error) {

	trace.Func()

	return u.postRepo.GetList(ctx, limit, offset, isPage, isPublished)
}

func (u PostUc) Create(ctx context.Context, title, content string, isPage, isPublished bool) (post models.Post, err error) {

	trace.Func()

	err = u.AuthorizationUc.AuthorizeUser(ctx)
	if err != nil {
		return
	}

	post = models.Post{
		Title:       title,
		Content:     content,
		IsPublished: isPublished,
		IsPage:      isPage,
	}

	err = u.postRepo.Create(ctx, &post)
	if err != nil {
		post = models.Post{}
	}

	return
}

func (u PostUc) Update(ctx context.Context,
	id int, slug, title, content string, isPublished bool) (post models.Post, err error) {

	trace.Func()

	err = u.AuthorizationUc.AuthorizeUser(ctx)
	if err != nil {
		return
	}
	postInDB, err := u.postRepo.GetById(ctx, id, nil)
	if err != nil {
		return
	}

	if title == "" {
		title = postInDB.Title
	}

	if slug == "" {
		slug = postInDB.Slug
	}

	if content == "" {
		content = postInDB.Content
	}

	post = models.Post{
		Id:          id,
		Slug:        slug,
		Title:       title,
		Content:     content,
		IsPublished: isPublished,
		PublishedAt: postInDB.PublishedAt,
	}

	err = u.postRepo.Update(ctx, &post)

	return post, err
}

func (u PostUc) Delete(ctx context.Context, slug string) (err error) {

	trace.Func()

	err = u.AuthorizationUc.AuthorizeUser(ctx)
	if err != nil {
		return
	}
	err = u.postRepo.Delete(ctx, slug)

	return
}
