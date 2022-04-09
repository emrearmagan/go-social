/*
reddit.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package reddit

import (
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth/oauth2"
)

const (
	Base                = "https://oauth.reddit.com"
	AuthorizationPrefix = "bearer " // trailing space is required

	UserAgentHeaderKey = "User-Agent"
)

type Client struct {
	Account *AccountService
	User    *UserService
}

// NewClient returns a new Reddit Client.
func NewClient(oauth *oauth2.OAuth2, userAgent string) *Client {
	client := oauth.Client().New().Base(Base)
	client.Set(UserAgentHeaderKey, userAgent)
	oauth = oauth.NewClient(client)
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		Account: newAccountService(oauth),
		User:    newUserService(oauth),
	}
}

func (r *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := r.User.UserCredentials()
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
