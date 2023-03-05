package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/usecases"
)

const (
	ObjectNamePost = "Post"

	PostQueryGet       = "GetPost"
	PostQueryList      = "ListPosts"
	PostMutationCreate = "CreatePost"
	PostMutationUpdate = "UpdatePost"
	PostMutationDelete = "DeletePost"
)

const (
	ArgsPostsTitle       = "title"
	ArgsPostsContent     = "content"
	ArgsPostsId          = "id"
	ArgsPostsSlug        = "slug"
	ArgsPostsIsPublished = "isPublished"
	ArgsPostsIsPage      = "isPage"
)

type PostCtrInterface interface {
	GetBySlugFromGQL(ctx context.Context, args map[string]interface{}) (post models.Post, err error)
	GetListFromGQL(ctx context.Context, args map[string]interface{}) (posts []models.Post, err error)

	CreateFromGQL(ctx context.Context, args map[string]interface{}) (post models.Post, err error)

	UpdateFromGQL(ctx context.Context, args map[string]interface{}) (post models.Post, err error)

	DeleteFromGQL(ctx context.Context, args map[string]interface{}) (err error)
}

type PostCtr struct {
	postUc usecases.PostUcInterface
}

func NewPostCtr(postUc usecases.PostUcInterface) PostCtrInterface {
	return &PostCtr{postUc: postUc}
}
