/*
user.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitter

import (
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social"
)

const (
	UserPath = "/1.1/account/verify_credentials.json"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth1 *oauth1.OAuth1
}

// newUserService returns a new Twitter UserService.
func newUserService(oauth1 *oauth1.OAuth1) *UserService {
	return &UserService{
		oauth1: oauth1,
	}
}

// UserCredentials returns the authorized user if credentials are valid and returns an error otherwise.
// https://dev.twitter.com/rest/reference/get/account/verify_credentials
func (u *UserService) UserCredentials(params *UserCredentialsParams) (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := u.oauth1.Get(UserPath, user, apiError, params)
	return user, social.CheckError(err)
}

// UserCredentialsParams are the params for UserCredentials.
type UserCredentialsParams struct {
	// The entities node will not be included when set to false .
	IncludeEntities bool `url:"include_entities,omitempty"`
	// When set to either true , t or 1 statuses will not be
	// included in the returned user object.
	SkipStatus bool `url:"skip_status,omitempty"`
	// When set to true email will be returned in the user
	// objects as a string. If the user does not have an email
	// address on their account, or if the email address is not verified,
	// null will be returned.
	IncludeEmail bool `url:"include_email,omitempty"`
}

type User struct {
	ID                             int64       `json:"id"`
	Name                           string      `json:"name"`
	ScreenName                     string      `json:"screen_name"`
	FollowersCount                 int         `json:"followers_count"`
	FriendsCount                   int         `json:"friends_count"`
	CreatedAt                      string      `json:"created_at"`
	Verified                       bool        `json:"verified"`
	StatusCount                    int         `json:"statuses_count"`
	ProfileBackgroundImageURL      interface{} `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS interface{} `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
	Suspended                      bool        `json:"suspended"`
}
