/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package models

import "errors"

var (
	// Business layer
	ErrBadRequst             = errors.New("bad request")
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
