/*
user.go
Created at 09.07.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package youtube

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"github.com/emrearmagan/go-social/social/client"
)

const (
	UserBase = "https://www.googleapis.com/"
	UserPath = "/oauth2/v3/userinfo"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new YouTube UserService.
func newUserService(oauth2 *oauth2.OAuth2) *UserService {
	return &UserService{
		oauth2: oauth2,
	}
}

// UserInfo returns the authorized user if credentials are valid and returns an error otherwise.
// Required scopes: https://www.googleapis.com/auth/userinfo.profile
func (u *UserService) UserInfo() (*UserInfoResp, error) {
	user := new(UserInfoResp)
	apiError := new(APIError)

	// Requires a different base
	cl := client.NewHttpClient().Base(UserBase)
	auther := u.oauth2.NewClient(cl)
	err := auther.Get(UserPath, user, apiError, nil)
	return user, social.CheckError(err)
}

type UserInfoResp struct {
	Sub        string `json:"sub"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}
