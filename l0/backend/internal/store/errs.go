package store

import "errors"

var (
	ErrWriteCache = errors.New("write error to table orders")
	ErrNotFound = errors.New("error, item not found")
)