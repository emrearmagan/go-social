/*
errors.go
Created at 09.07.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package youtube

import (
	"fmt"
	"github.com/emrearmagan/go-social/models/errors"
)

// APIError represents a Twitch API error with its corresponding http StatusCode response
// https://dev.twitch.tv/docs/api/
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents the actual error response from the Api
type ErrorDetail struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
		Status string `json:"status"`
	} `json:"error"`
}

func (e *APIError) ErrorDetail() interface{} {
	return &e.Errors
}

func (e *APIError) Error() string {
	if len(e.Errors.Error.Message) > 0 {
		return fmt.Sprintf("youtube: %d - %v", e.StatusCode, e.Errors)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *APIError) Empty() bool {
	if len(e.Errors.Error.Message) == 0 {
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
	case 400:
		return errors.New(errors.ErrBadRequest, e.Error())
	case 401, 403: // Unauthenticated: Missing/invalid Token or missing scopes
		return errors.New(errors.ErrUnauthorized, e.Error())
	case 404:
		return errors.New(errors.ErrNotFound, e.Error())
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return errors.New(errors.ErrRateLimit, e.Error())
	case 500, 502, 503: //Internal api error
		return errors.New(errors.ErrApiError, e.Error())

	}
	return errors.New(errors.ErrUnknownError, e.Error())
}
