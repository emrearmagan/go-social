/*
user.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"go-social/social"
	"go-social/social/oauth/oauth2"
	"time"
)

const (
	UserPath = "/v2/user/"
)

// UserService provides methods for user credentials
type UserService struct {
	oauth2 *oauth2.OAuth2
}

// newUserService returns a new Github UserService.
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

// User represents a Dribble User.
// https://developer.dribbble.com/v2/user/
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Links     struct {
		Web     string `json:"web"`
		Twitter string `json:"twitter"`
	} `json:"links"`
	CanUploadShot  bool      `json:"can_upload_shot"`
	Pro            bool      `json:"pro"`
	FollowersCount int       `json:"followers_count"`
	CreatedAt      time.Time `json:"created_at"`
	Type           string    `json:"type"`
	Teams          []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Login     string `json:"login"`
		HTMLURL   string `json:"html_url"`
		AvatarURL string `json:"avatar_url"`
		Bio       string `json:"bio"`
		Location  string `json:"location"`
		Links     struct {
			Web     string `json:"web"`
			Twitter string `json:"twitter"`
		} `json:"links"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"teams"`
}
