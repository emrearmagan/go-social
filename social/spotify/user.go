/*
user.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
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
