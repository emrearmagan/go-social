/*
user.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package github

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
)

const (
	UserPath = "/user"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new GitHub UserService.
func newUserService(oauth2 *oauth2.OAuth2) *UserService {
	return &UserService{
		oauth2: oauth2,
	}
}

// UserCredentials returns the user credentials for the authenticated user.
// https://developer.dribbble.com/v2/user/
func (u *UserService) UserCredentials() (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := u.oauth2.Get(UserPath, user, apiError, nil)
	return user, social.CheckError(err)
}
