/*
user.go
Created at 19.06.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitch

import (
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"time"
)

const (
	UserPath = "/helix/users/"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new Twitch UserService.
func newUserService(oauth2 *oauth2.OAuth2) *UserService {
	return &UserService{
		oauth2: oauth2,
	}
}

// UserCredentials returns the authorized user if credentials are valid and returns an error otherwise.
// https://dev.twitch.tv/docs/api/reference#get-users
// Required scopes: -
func (u *UserService) UserCredentials(params *UserParams) (*User, error) {
	user := new(User)
	apiError := new(APIError)

	err := u.oauth2.Get(UserPath, user, apiError, params)
	return user, social.CheckError(err)
}

// UserParams are the params for UserCredentials.
type UserParams struct {
	// ID. Multiple user IDs can be specified. Limit: 100. Optional
	ID string `url:"id,omitempty"`
	// Login name. Multiple login names can be specified. Limit: 100. Optional
	Login string `url:"login,omitempty"`
}

type User struct {
	Data []struct {
		ID              string    `json:"id"`
		Login           string    `json:"login"`
		DisplayName     string    `json:"display_name"`
		Type            string    `json:"type"`
		BroadcasterType string    `json:"broadcaster_type"`
		Description     string    `json:"description"`
		ProfileImageURL string    `json:"profile_image_url"`
		OfflineImageURL string    `json:"offline_image_url"`
		ViewCount       int       `json:"view_count"`
		Email           string    `json:"email"`
		CreatedAt       time.Time `json:"created_at"`
	} `json:"data"`
}
