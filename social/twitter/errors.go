/*
errors.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitter

import (
	"fmt"
	"github.com/emrearmagan/go-social/models/errors"
)

// APIError represents a Spotify API error with its corresponding http StatusCode response
// https://developer.twitter.com/ja/docs/basics/response-codes
type APIError struct {
	StatusCode int
	Errors     ErrorDetail
}

// ErrorDetail represents the actual error response from the Api
type ErrorDetail struct {
	ErrorStruct []ErrorStruct `json:"error"`
}

type ErrorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) ErrorDetail() interface{} {
	return &e.Errors
}

func (e *APIError) Error() string {
	if len(e.Errors.ErrorStruct) > 0 {
		err := e.Errors.ErrorStruct[0]
		return fmt.Sprintf("twitter: %d - %v", err.Code, err.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e *APIError) Empty() bool {
	if len(e.Errors.ErrorStruct) == 0 {
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

// TODO: error handling for not modified
func (e *APIError) ReturnErrorResponse() error {
	switch e.Status() {
	case int(invalidCoordinates), int(parameterMissing):
		return errors.New(errors.ErrBadRequest, e.Error())
	case int(rateLimitExceeded):
		return errors.New(errors.ErrRateLimit, e.Error())
	case int(invalidExpiredToken):
		return errors.New(errors.ErrInvalidOrExpiredToken, e.Error())
	case int(internalError):
		return errors.New(errors.ErrApiError, e.Error())
	case int(badAuthenticationData):
		return errors.New(errors.ErrBadAuthenticationData, e.Error())
	case int(endpointRetired), int(pageDoesNotExists), int(userNotFound), int(noLocation), int(noUserMatches):
		return errors.New(errors.ErrNotFound, e.Error())
	case int(userSuspended), int(accountSuspended), int(appNotAllowedToAccessOrDeleteMessage), int(credentialsDontAllowThisResource), int(appCannotPerformWriteActions):
		return errors.New(errors.ErrForbidden, e.Error())
	case int(invalidSuspendedApp), int(couldNotAuthenticate), int(notPermitted), int(unableToVerifyCreds), int(notAuthorizedForThisStatus):
		return errors.New(errors.ErrUnauthorized, e.Error())
	}

	return errors.New(errors.ErrUnknownError, e.Error())
}

// Most of the error codes from Twitter API.
// See: https://developer.twitter.com/en/support/twitter-api/error-troubleshooting for more
type twitterApiCode int

const (
	invalidCoordinates    twitterApiCode = 3   // Corresponds with HTTP 400. The coordinates provided as parameters were not valid for the request.
	badAuthenticationData twitterApiCode = 215 // Corresponds with HTTP 400. The method requires authentication but it was not presented or was wholly invalid.

	couldNotAuthenticate twitterApiCode = 32  // Corresponds with HTTP 401. There was an issue with the authentication data for the request.
	invalidSuspendedApp  twitterApiCode = 416 // Corresponds with HTTP 401. The App has been suspended and cannot be used with Sign-in with Twitter.

	parameterMissing                     twitterApiCode = 38  // Corresponds with HTTP 403. The request is missing the <named> parameter (such as media, text, etc.) in the request.
	userSuspended                        twitterApiCode = 63  // Corresponds with HTTP 403 The user account has been suspended and information cannot be retrieved.
	accountSuspended                     twitterApiCode = 64  // Corresponds with HTTP 403. The access token being used belongs to a suspended user.
	notPermitted                         twitterApiCode = 87  //Corresponds with HTTP 403. The endpoint called is not a permitted URL.
	invalidExpiredToken                  twitterApiCode = 89  // Corresponds with HTTP 403. The access token used in the request is incorrect or has expired.
	appNotAllowedToAccessOrDeleteMessage twitterApiCode = 93  //Corresponds with HTTP 403. The OAuth token does not provide access to Direct Messages.
	unableToVerifyCreds                  twitterApiCode = 99  // Corresponds with HTTP 403. The OAuth credentials cannot be validated. Check that the token is still valid.
	credentialsDontAllowThisResource     twitterApiCode = 220 //Corresponds with HTTP 403. The authentication token in use is restricted and cannot access the requested resource.
	notAuthorizedForThisStatus           twitterApiCode = 179 //Corresponds with HTTP 403. Thrown when a Tweet cannot be viewed by the authenticating user, usually due to the Tweet’s author having protected their Tweets.
	appCannotPerformWriteActions         twitterApiCode = 261 // Corresponds with HTTP 403. Caused by the App being restricted from POST, PUT, or DELETE actions.

	pageDoesNotExists twitterApiCode = 34 // Corresponds with HTTP 404. The specified resource was not found.
	userNotFound      twitterApiCode = 50 // Corresponds with HTTP 404. The user is not found.
	noLocation        twitterApiCode = 13 // Corresponds with HTTP 404. It was not possible to derive a location for the IP address provided as a parameter on the geo search request.
	noUserMatches     twitterApiCode = 17 // Corresponds with HTTP 404. It was not possible to find a user profile matching the parameters specified.

	endpointRetired   twitterApiCode = 251 // Corresponds with HTTP 410. The App made a request to a retired URL.
	rateLimitExceeded twitterApiCode = 88  //Corresponds with HTTP 429. The request limit for this resource has been reached for the current rate limit window.
	internalError     twitterApiCode = 131 //Corresponds with HTTP 500. An unknown internal error occurred.

)
