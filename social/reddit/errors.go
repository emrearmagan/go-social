/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"fmt"
	"go-social/models"
)

// APIError represents a Reddit API StatusCode response
// https://developer.spotify.com/documentation/web-api/
type APIError struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error"`
}

func (e APIError) Error() string {
	if (e != APIError{}) {
		return fmt.Sprintf("Reddit: %d - %v", e.ErrorCode, e.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e APIError) Empty() bool {
	if (e == APIError{}) {
		return true
	}
	return false
}

func (e APIError) Status() int {
	return e.ErrorCode
}

func (e APIError) ReturnErrorResponse() error {
	switch e.Status() {
	case 304: // The content has not been modified and client should use cached data
		return models.ErrNotModified
	case 401, 403: // Invalid or expired token (if token has been revoked) - The access token used in the request is incorrect or has expired.
		return models.ErrUnauthorized
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return models.ErrRateLimit
	case 500, 502, 503: //Internal api error
		return models.ErrApiError
	}
	return models.ErrUnknownError
}
