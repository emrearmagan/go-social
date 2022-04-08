/*
user.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
)

const (
	UserPath = "/v2/user/"
)

// UserService provides a method for account credential verification.
type UserService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new AccountService.
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

	err := get(u.oauth2, UserPath, user, apiError, nil)

	return user, social.CheckError(err)
}
