package tonclient

import "errors"

var (
	ErrFailedToCreateRequest = errors.New("failed to create request")
	ErrFailedToParseAddress  = errors.New("failed to parse address")
	ErrFailedAtDuringRequest = errors.New("failed at during request")
	ErrBadRequest            = errors.New("bad request")
)
