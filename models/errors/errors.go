/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package errors

import (
	"errors"
)

var (
	ErrBadRequest            = errors.New("bad request")
	ErrNotFound              = errors.New("not found")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrRateLimit             = errors.New("rate limit exceeded")
	ErrBadAuthenticationData = errors.New("bad authentication data")
	ErrInvalidOrExpiredToken = errors.New("invalid or expired token")
	ErrForbidden             = errors.New("user/app has been suspended or deleted")
	ErrNotModified           = errors.New("not modified")
	ErrApiError              = errors.New("internal api error")

	ErrUnknownError = errors.New("unknown error")
)

type SocialError struct {
	Errors  error
	Message string
}

func New(error error, message string) SocialError {
	return SocialError{
		Errors:  error,
		Message: message,
	}
}

func (s SocialError) Error() string {
	return s.Message
}
