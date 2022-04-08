/*
dribbble.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/social"
	"strings"
)

const (
	APIHost   = "https://api.dribbble.com/"
	UserPath  = "/v2/user/"
	ShotsPath = "/v2/user/shots"

	AuthorizationHeaderName = "Authorization"
	authorizationPrefix     = "Bearer " // trailing space is required
)

func get(token string, path string, resp interface{}, apiError *APIError) error {
	client := social.NewClient().Base(APIHost).Get(path)

	req, err := client.Request()
	if err != nil {
		return err
	}

	header := []string{authorizationPrefix, token}
	req.Header.Set(AuthorizationHeaderName, strings.Join(header, ""))

	httpResp, err := client.Do(req, resp, &apiError.Errors)
	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.StatusCode = httpResp.StatusCode
		}
	}

	return social.RelevantError(err, apiError)
}

// DribbbleShots returns all shots for the authenticated user.
// See: https://developer.dribbble.com/v2/shots/#list-shots for more information
func DribbbleShots(token string) (*Shots, error) {
	shots := new(Shots)
	apiError := new(APIError)

	err := get(token, ShotsPath, &shots, apiError)
	if err != nil {
		return nil, err
	}

	return shots, social.CheckError(err)
}

// UserCredentials returns the user credentials for the authenticated user.
// https://developer.dribbble.com/v2/user/
func UserCredentials(token string) (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := get(token, UserPath, &user, apiError)
	if err != nil {
		return nil, err
	}

	return user, social.CheckError(err)
}
