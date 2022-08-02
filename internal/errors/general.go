package errors

import "net/http"

var (
	// ErrParseRequest General error for parse request payload
	ErrParseRequest = New(http.StatusBadRequest, "unable processing request")
)
