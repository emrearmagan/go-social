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
	Account *AccountService
	User    *UserService
}

// NewClient returns a new Dribbble Client.
func NewClient(oauth *oauth2.OAuth2, userAgent string) *Client {
	//TODO: Create client here and add the header before adding it to the oauth.
	// That way i dont need to put the user agent every where
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		Account: newAccountService(oauth, userAgent),
		User:    newUserService(oauth, userAgent),
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
