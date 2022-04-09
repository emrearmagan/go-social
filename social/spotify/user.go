/*
user.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
)

const (
	UserPath = "/v1/me/"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new Spotify UserService.
func newUserService(oauth2 *oauth2.OAuth2) *UserService {
	return &UserService{
		oauth2: oauth2,
	}
}

// UserCredentials returns the user credentials for the authenticated user.
// https://developer.spotify.com/console/get-current-user/
func (u *UserService) UserCredentials() (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := u.oauth2.Get(UserPath, user, apiError, nil)
	return user, social.CheckError(err)
}

type User struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		URL    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Product string `json:"product"`
	Type    string `json:"type"`
	URI     string `json:"uri"`
}
