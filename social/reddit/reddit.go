/*
reddit.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"go-social/models"
	"go-social/social/oauth/oauth2"
)

const (
	Base                = "https://oauth.reddit.com"
	AuthorizationPrefix = "bearer " // trailing space is required
)

type Client struct {
	User *UserService
}

// NewClient returns a new Dribbble Client.
func NewClient(oauth *oauth2.OAuth2, userAgent string) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		User: newUserService(oauth.New(), userAgent),
	}
}

func (d *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := d.User.UserCredentials()
	if err != nil {
		return nil, err
	}

	goSocial := models.SocialUser{
		Username:     u.Name,
		Name:         u.Subreddit.DisplayName,
		UserId:       u.ID,
		Verified:     u.Verified,
		ContentCount: int64(u.TotalKarma),
		Following:    &u.NumFriends,
		AvatarUrl:    u.SnoovatarImg,
		Followers:    u.Subreddit.Subscribers,
		Url:          "https://www.reddit.com" + u.Subreddit.URL,
	}

	return &goSocial, nil
}
