package controllers

import "errors"

var (
	errInputInvalid = errors.New("input is invalid")
)

var (
	ArgsId          = "id"
	ArgsCreatedAt   = "createdAt"
	ArgsUpdatedAt   = "updatedAt"
	ArgsPublishedAt = "publishedAt"

	ArgsLimit = "limit"
	ArgsPage  = "page"
)
