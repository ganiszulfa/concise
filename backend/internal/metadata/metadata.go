package metadata

import (
	"context"
	"errors"

	"github.com/ganiszulfa/concise/backend/config/app"
	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
)

var errMsgInputInvalid = "input is invalid"

func GetById(ctx context.Context, args map[string]interface{}) (md models.Metadata, err error) {
	trace.Func()

	id, ok := args["id"].(int)
	if !ok {
		return md, errors.New(errMsgInputInvalid)
	}

	result := app.DB.WithContext(ctx).First(&md, "id = ?", id)
	return md, result.Error
}

func GetByKey(ctx context.Context, args map[string]interface{}) (md models.Metadata, err error) {
	trace.Func()

	key, ok := args["key"].(string)
	if !ok {
		return md, errors.New(errMsgInputInvalid)
	}

	result := app.DB.WithContext(ctx).First(&md, "key = ?", key)
	return md, result.Error
}

func GetList(ctx context.Context, args map[string]interface{}) (mds []models.Metadata, err error) {
	trace.Func()

	page, ok := args["page"].(int)
	if !ok || page == 0 {
		page = 1
	}

	limit, ok := args["limit"].(int)
	if !ok || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	result := app.DB.WithContext(ctx).
		Limit(limit).Offset(offset).Find(&mds)

	return mds, result.Error
}

func Create(ctx context.Context, args map[string]interface{}) (md models.Metadata, err error) {
	trace.Func()

	key, ok := args["key"].(string)
	if !ok {
		return md, errors.New(errMsgInputInvalid)
	}

	value, ok := args["value"].(string)
	if !ok {
		return md, errors.New(errMsgInputInvalid)
	}

	md = models.Metadata{
		Key:   key,
		Value: value,
	}

	result := app.DB.Create(&md)

	return md, result.Error
}

func Update(ctx context.Context, args map[string]interface{}) (md models.Metadata, err error) {
	trace.Func()

	md, err = GetByKey(ctx, args)
	if err != nil {
		return md, err
	}

	value, ok := args["value"].(string)
	if !ok {
		md.Value = value
	}

	return models.Metadata{}, nil
}
