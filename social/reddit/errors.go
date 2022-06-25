/*
errors.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"fmt"
	"github.com/emrearmagan/go-social/models/errors"
)

// APIError represents a Reddit API error with its corresponding http StatusCode response
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

type ErrorDetail struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

func (e *APIError) ErrorDetail() interface{} {
	return &e.Errors
}

func (e *APIError) Error() string {
	if (e.Errors != ErrorDetail{}) {
		return fmt.Sprintf("Reddit: %d - %v : %v", e.StatusCode, e.Errors.Error, e.Errors.Message)
	}
	return fmt.Sprintf("Reddit: %d", e.StatusCode)
}

// Empty returns true if nil. Otherwise, at least 1 error message/code is
func (e *APIError) Empty() bool {
	return e.Errors == ErrorDetail{}
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
		return errors.New(errors.ErrNotModified, e.Error())
	case 401, 403: // Invalid or expired token (if token has been revoked) - The access token used in the request is incorrect or has expired.
		return errors.New(errors.ErrUnauthorized, e.Error())
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return errors.New(errors.ErrRateLimit, e.Error())
	case 400, 411:
		return errors.New(errors.ErrBadRequest, e.Error())
	case 500, 502, 503: //Internal api error
		return errors.New(errors.ErrApiError, e.Error())
	}
	return errors.New(errors.ErrUnknownError, e.Error())
}
