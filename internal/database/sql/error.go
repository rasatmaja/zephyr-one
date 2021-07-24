package sql

import (
	"errors"
)

var (
	// ErrNotFound define error for empty rows result
	ErrNotFound = errors.New("DATA_NOT_FOUND")
	// ErrDataDuplicate define error for data duplication
	ErrDataDuplicate = errors.New("DATA_DUPLICATE")
)
