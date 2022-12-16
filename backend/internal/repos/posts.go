package repos

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type PostRepoInterface interface {
	GetBySlug(ctx context.Context, slug string, isPublished *bool) (post models.Post, err error)
	GetList(ctx context.Context, limit, offset int, isPage, isPublished *bool) (posts []models.Post, err error)

	Create(ctx context.Context, post *models.Post) (err error)
	Update(ctx context.Context, post *models.Post) (err error)
	Delete(ctx context.Context, slug string) (err error)
}

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepoInterface {
	return &PostRepo{db: db}
}

func (r PostRepo) GetBySlug(ctx context.Context, slug string, isPublished *bool) (post models.Post, err error) {
	trace.Func()

	result := r.db.WithContext(ctx)

	if isPublished != nil {
		result = result.
			Where("is_published", isPublished)
	}

	result = result.
		Where("is_deleted", false)

	result = result.First(&post, "slug = ?", slug)

	return post, result.Error
}

func (r PostRepo) GetList(ctx context.Context, limit, offset int, isPage, isPublished *bool) (posts []models.Post, err error) {
	trace.Func()

	result := r.db.WithContext(ctx).
		Limit(limit).Offset(offset).Order(`"created_at" desc`)

	if isPublished != nil {
		result = result.
			Where("is_published", isPublished)
	}

	if isPage != nil {
		result = result.
			Where("is_page", isPage)
	}

	result = result.
		Where("is_deleted", false)

	err = result.Find(&posts).Error

	return
}

func (r PostRepo) Create(ctx context.Context, post *models.Post) (err error) {
	trace.Func()

	post.Slug = slug.Make(post.Title)

	post.CreatedAt = time.Now()

	if post.IsPublished {
		post.PublishedAt = time.Now()
	}

	err = r.db.Create(post).Error

	if err != nil {
		// retry again
		post.Slug = generateSafePostSlug(post.Title)
		err = r.db.Create(post).Error
	}

	return
}

func (r PostRepo) Update(ctx context.Context, post *models.Post) (err error) {

	trace.Func()

	post.UpdatedAt = time.Now()

	if post.IsPublished && post.PublishedAt.IsZero() {
		post.PublishedAt = post.UpdatedAt
	}

	err = app.DB.Model(post).
		Where("slug = ?", post.Slug).
		Updates(
			map[string]interface{}{
				"title":        post.Title,
				"content":      post.Content,
				"is_published": post.IsPublished,
				"updated_at":   post.UpdatedAt,
				"published_at": post.PublishedAt,
			},
		).Error

	return
}

func (r PostRepo) Delete(ctx context.Context, slug string) (err error) {
	trace.Func()

	post := models.Post{}
	err = app.DB.Model(&post).
		Where("slug = ?", slug).
		Updates(
			map[string]interface{}{
				"is_deleted": true,
				"updated_at": time.Now(),
			},
		).Error

	return
}

func generateSafePostSlug(s string) string {
	i := 10 * 1000
	r := fmt.Sprintf("_%d", rand.Intn(i*10)+i)
	return slug.Make(s) + r
}
