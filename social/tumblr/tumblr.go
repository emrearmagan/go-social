/*
tumblr.go
Created at 10.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package tumblr

import (
	"context"
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth1"
	"github.com/emrearmagan/go-social/social/client"
)

const (
	Base = "https://api.tumblr.com/"
)

type Client struct {
	User *UserService
}

// NewClient returns a new Spotify Client.
func NewClient(ctx context.Context, c *oauth.Credentials, token *oauth1.Token) *Client {
	cl := client.NewHttpClient().Base(Base)
	auther := oauth1.NewOAuth(ctx, c, token, cl)
	return &Client{
		User: newUserService(auther),
	}
}

func (s *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := s.User.UserCredentials()
	if err != nil {
		return nil, err
	}

	follower := 0
	posts := int64(0)
	for _, b := range u.Response.User.Blogs {
		follower += b.Followers
		posts += int64(b.Posts)
	}

	goSocial := models.SocialUser{
		Username: u.Response.User.Name,
		Name:     u.Response.User.Name,
		//Does not provide a user id ?
		UserId:       u.Response.User.Name,
		ContentCount: posts,
		Verified:     false,
		AvatarUrl:    u.Response.User.Blogs[0].Avatar[0].URL,
		Following:    &u.Response.User.Following,
		Followers:    follower,
		Url:          u.Response.User.Blogs[0].URL,
	}

	return &goSocial, nil
}
