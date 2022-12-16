package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

func (c MetadataCtr) GetAllFromGQL(ctx context.Context, args map[string]interface{}) (mds []models.Metadata, err error) {
	trace.Func()

	return c.metadataUc.GetAll(ctx)
}
