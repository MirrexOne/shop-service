package v1

import "fmt"

var (
	ErrInvalidAuthHeader = fmt.Errorf("invalid auth header")
	ErrCannotParseToken  = fmt.Errorf("cannot parse token")
)

type errorResponse struct {
	errStatus int
	message   string
}

func newErrorResponse(errStatus int, message string) *errorResponse {
	return &errorResponse{
		errStatus: errStatus,
		message:   message,
	}
}
