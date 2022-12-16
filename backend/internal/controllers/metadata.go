package controllers

import (
	"context"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/usecases"
)

const (
	ObjectNameMetadata = "Metadata"

	MetadataQueryList = "ListMetadata"
)

const (
	ArgsMetadataKey   = "key"
	ArgsMetadataValue = "value"
)

type MetadataCtrInterface interface {
	GetAllFromGQL(ctx context.Context, args map[string]interface{}) (mds []models.Metadata, err error)
}

type MetadataCtr struct {
	metadataUc usecases.MetadataUcInterface
}

func NewMetadataCtr(metadataUc usecases.MetadataUcInterface) MetadataCtrInterface {
	return &MetadataCtr{metadataUc: metadataUc}
}
