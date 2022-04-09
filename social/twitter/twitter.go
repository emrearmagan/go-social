/*
twitter.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitter

import (
	"fmt"
	"go-social/models"
	"go-social/social/oauth/oauth1"
)

type Client struct {
	User *UserService
}

const (
	Base = "https://api.twitter.com/"
)

// NewClient returns a new Twitter Client.
func NewClient(oauth *oauth1.OAuth1) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	return &Client{
		User: newUserService(oauth),
	}
}

func (r *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := r.User.UserCredentials(nil)
	if err != nil {
		return nil, err
	}

	goSocial := models.SocialUser{
		Username:     u.ScreenName,
		Name:         u.Name,
		UserId:       fmt.Sprintf("%v", u.ID),
		Verified:     u.Verified,
		ContentCount: int64(u.StatusCount),
		AvatarUrl:    u.ProfileImageURL,
		Followers:    u.FollowersCount,
		Following:    &u.FriendsCount,
		Url:          fmt.Sprintf("https://twitter.com/%s", u.ScreenName),
	}

	return &goSocial, nil
}
