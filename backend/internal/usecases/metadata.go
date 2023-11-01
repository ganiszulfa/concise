package usecases

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/models/keys"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

type MetadataUcInterface interface {
	GetAll(ctx context.Context) (mds []models.Metadata, err error)
	GetByKey(ctx context.Context, key string) (md models.Metadata, err error)
}

type MetadataUc struct {
	metadataRepo repos.MetadataRepoInterface
}

func NewMetadataUc(metadataRepo repos.MetadataRepoInterface) MetadataUcInterface {
	return &MetadataUc{metadataRepo: metadataRepo}
}

func (u MetadataUc) GetAll(ctx context.Context) (mds []models.Metadata, err error) {

	trace.Func()

	mds, err = u.metadataRepo.GetAll(ctx)
	if err != nil {
		return
	}

	for i, v := range mds {
		if v.Key == keys.KEY_USER_PASSWORD {
			mds[i].Value = ""
		}
	}

	return
}

func (u MetadataUc) GetByKey(ctx context.Context, key string) (md models.Metadata, err error) {

	trace.Func()

	md, err = u.metadataRepo.GetByKey(ctx, key)
	return
}

func (u MetadataUc) UpdateByKey(ctx context.Context, key, value string) (md models.Metadata, err error) {

	trace.Func()

	md = models.Metadata{
		Key:   key,
		Value: value,
	}

	err = u.metadataRepo.UpdateByKey(ctx, md)
	return
}
