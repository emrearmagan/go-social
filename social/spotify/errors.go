/*
errors.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"fmt"
	"github.com/emrearmagan/go-social/models/errors"
)

// APIError represents a Spotify API error with its corresponding http StatusCode response
// https://developer.spotify.com/documentation/web-api/
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents the actual error response from the Api
type ErrorDetail struct {
	ErrorStruct ErrorStruct `json:"error"`
}

type ErrorStruct struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func (e *APIError) ErrorDetail() interface{} {
	return &e.Errors
}

func (e *APIError) Error() string {
	if (e.Errors != ErrorDetail{}) {
		return fmt.Sprintf("Spotify: %d - %v", e.Errors.ErrorStruct.StatusCode, e.Errors.ErrorStruct.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *APIError) Empty() bool {
	if (e.Errors == ErrorDetail{}) {
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
		return errors.New(errors.ErrNotModified, e.Error())
	case 401, 403: // Invalid or expired token (if token has been revoked) - The access token used in the request is incorrect or has expired.
		return errors.New(errors.ErrUnauthorized, e.Error())
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return errors.New(errors.ErrRateLimit, e.Error())
	case 500, 502, 503: //Internal api error
		return errors.New(errors.ErrApiError, e.Error())
	}
	return errors.New(errors.ErrUnknownError, e.Error())
}

//RefreshError
type RefreshError struct {
	StatusCode int
	Errors     RefreshDetail
}

type RefreshDetail struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e *RefreshError) ErrorDetail() interface{} {
	return &e.Errors
}

func (e *RefreshError) Error() string {
	if (e.Errors != RefreshDetail{}) {
		return fmt.Sprintf("Spotify: %d - %v : %v", e.StatusCode, e.Errors.Error, e.Errors.ErrorDescription)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *RefreshError) Empty() bool {
	return e.Errors == RefreshDetail{}
}

func (e *RefreshError) Status() int {
	return e.StatusCode
}

func (e *RefreshError) SetStatus(code int) {
	e.StatusCode = code
}

func (e *RefreshError) ReturnErrorResponse() error {
	switch e.Status() {
	case 400:
		return errors.New(errors.ErrBadRequest, e.Error())
	case 401, 403: // Invalid or expired token (if token has been revoked) - The access token used in the request is incorrect or has expired.
		return errors.New(errors.ErrUnauthorized, e.Error())
	case 429: // Rate limit exceeded	 - The request limit for this resource has been reached for the current rate limit window.
		return errors.New(errors.ErrRateLimit, e.Error())
	case 500, 502, 503: //Internal api error
		return errors.New(errors.ErrApiError, e.Error())
	}
	return errors.New(errors.ErrUnknownError, e.Error())
}
