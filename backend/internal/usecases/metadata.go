package usecases

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/repos"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

type MetadataUcInterface interface {
	GetAll(ctx context.Context) (mds []models.Metadata, err error)
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

	return
}
