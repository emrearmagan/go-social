/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"fmt"
	"github.com/emrearmagan/go-social/models"
)

//TODO: Reddit errors are not properly handled. Sometimes instead of an error html is returned

// APIError represents a Reddit API error with its corresponding http StatusCode response
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents the actual error response from the Api
type ErrorDetail struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error"`
}

func (e *APIError) ErrorDetail() interface{} {
	return &e.Errors
}

func (e *APIError) Error() string {
	if len(e.Errors.Message) > 0 {
		return fmt.Sprintf("Reddit: %d - %v", e.Errors.ErrorCode, e.Errors)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *APIError) Empty() bool {
	if len(e.Errors.Message) == 0 {
		return true
	}
	return false
}

func (e *APIError) Status() int {
	return e.StatusCode
}

func (e *APIError) SetStatus(code int) {
	e.StatusCode = code
}

func (e *APIError) ReturnErrorResponse() error {
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
