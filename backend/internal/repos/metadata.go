package repos

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"gorm.io/gorm"
)

type MetadataRepoInterface interface {
	GetAll(ctx context.Context) (mds []models.Metadata, err error)
	GetByKey(ctx context.Context, key string) (md models.Metadata, err error)
}

type MetadataRepo struct {
	db *gorm.DB
}

func NewMetadataRepo(db *gorm.DB) MetadataRepoInterface {
	return &MetadataRepo{db: db}
}

func (r MetadataRepo) GetAll(ctx context.Context) (mds []models.Metadata, err error) {
	trace.Func()

	err = r.db.WithContext(ctx).Limit(100).Offset(0).Find(&mds).Error

	return
}

func (r MetadataRepo) GetByKey(ctx context.Context, key string) (md models.Metadata, err error) {
	trace.Func()

	err = r.db.WithContext(ctx).First(&md, "key = ?", key).Error

	return
}
