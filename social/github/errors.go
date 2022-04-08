/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package github

import (
	"fmt"
	"go-social/models"
)

// APIError represents a Github API StatusCode response
// https://docs.github.com/en/free-pro-team@latest/developers/apps/identifying-and-authorizing-users-for-github-apps#handling-a-revoked-github-app-authorization
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents an individual item in an APIError.
type ErrorDetail struct {
	Message     string `json:"message"`
	Description string `json:"documentation_url"`
}

func (e APIError) Error() string {
	if len(e.Errors.Message) > 0 {
		return fmt.Sprintf("github: %d - %v", e.StatusCode, e.Errors)
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
	case 304: // The content has not been modified and client should use cached data
		return models.ErrNotModified
	case 401, 403: // Invalid or expired token or client is not permitted to perform this action.
		return models.ErrUnauthorized
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return models.ErrRateLimit
	case 500, 502, 503: //Internal api error
		return models.ErrApiError
	}

	return models.ErrUnknownError
}
