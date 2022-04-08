/*
user.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
)

const (
	UserPath = "/api/v1/me"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth2    *oauth2.OAuth2
	userAgent string
}

// newUserService returns a new Reddit UserService.
func newUserService(oauth2 *oauth2.OAuth2, userAgent string) *UserService {
	return &UserService{
		oauth2:    oauth2,
		userAgent: userAgent,
	}
}

// UserCredentials returns the user credentials for the authenticated user.
// https://www.reddit.com/dev/api#GET_api_v1_me
func (u *UserService) UserCredentials() (*User, error) {
	user := new(User)
	apiError := new(APIError)

	u.oauth2.AddCustomHeaders("User-Agent", u.userAgent)
	err := u.oauth2.Get(UserPath, user, apiError, nil)
	return user, social.CheckError(err)
}
