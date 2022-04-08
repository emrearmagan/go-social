/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"fmt"
	"go-social/models"
)

// APIError represents a Dribble API StatusCode response
// https://developer.dribbble.com/v2/#client-errors
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents an individual item in an APIError.
type ErrorDetail struct {
	Message string `json:"message"`
}

func (e APIError) Error() string {
	if len(e.Errors.Message) > 0 {
		return fmt.Sprintf("dribbble: %d - %v", e.StatusCode, e.Errors)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e APIError) Empty() bool {
	if len(e.Errors.Message) == 0 {
		return true
	}
	return false
}

func (e APIError) Status() int {
	return e.StatusCode
}

func (e APIError) ReturnErrorResponse() error {
	switch e.Status() {
	case 401, 403: // Invalid or expired token (if token has been revoked) - The access token used in the request is incorrect or has expired.
		return models.ErrUnauthorized
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return models.ErrRateLimit
	case 500, 502, 503: //Internal api error
		return models.ErrApiError

	}
	return models.ErrUnknownError
}
