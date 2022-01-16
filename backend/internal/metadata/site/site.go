package site

import (
	"context"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/internal/models/keys"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	log "github.com/sirupsen/logrus"
)

func getByKey(ctx context.Context, key string) (string, error) {
	trace.Func()

	md := models.Metadata{}
	result := app.DB.WithContext(ctx).First(&md, "key = ?", key)
	if result.Error != nil {
		return "", result.Error
	}
	return md.Value, nil
}

func GetOwnerPassword(ctx context.Context) string {
	trace.Func()

	v, err := getByKey(ctx, keys.KEY_OWNER_PASSWORD)
	if err != nil {
		log.Error(err)
		return "default"
	}

	return v
}
