/*
errors.go
Created at 10.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package tumblr

import (
	"fmt"
	"github.com/emrearmagan/go-social/models/errors"
)

// APIError represents a Tumblr API StatusCode response
// https://www.tumblr.com/docs/en/api/v2#common-response-elements
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents an individual item in an APIError.
type ErrorDetail struct {
	Meta struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	} `json:"meta"`
	Response []interface{} `json:"response"`
	Errors   []struct {
		Title  string `json:"title"`
		Code   int    `json:"code"`
		Detail string `json:"detail"`
	} `json:"errors"`
}

func (e *APIError) ErrorDetail() interface{} {
	return &e.Errors
}

//TODO: only returns the first error
func (e *APIError) Error() string {
	if len(e.Errors.Errors) > 0 {
		err := e.Errors.Errors[0]
		return fmt.Sprintf("Tumblr: %d - %v", err.Code, err.Detail)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *APIError) Empty() bool {
	if len(e.Errors.Errors) == 0 {
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

//TODO common error response, couldnt find propper documentation for error codes
func (e *APIError) ReturnErrorResponse() error {
	switch e.Status() {
	case 200:
		return nil
	case 304: // The content has not been modified and client should use cached data
		return errors.New(errors.ErrNotModified, e.Error())
	case 401, 403: // Invalid or expired token or client is not permitted to perform this action.
		return errors.New(errors.ErrUnauthorized, e.Error())
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return errors.New(errors.ErrRateLimit, e.Error())
	case 500, 502, 503: //Internal api error
		return errors.New(errors.ErrApiError, e.Error())
	}

	return errors.New(errors.ErrUnknownError, e.Error())
}
