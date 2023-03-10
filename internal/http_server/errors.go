package httpserver

import "errors"

var (
	ErrInvalidUserAddress       = errors.New("invalid user address")
	ErrInvalidCollectionAddress = errors.New("invalid collection address")
	ErrInvalidLimit             = errors.New("invalid limit")
	ErrInvalidOffset            = errors.New("invalid offset")
)
