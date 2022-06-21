/*
dribbble.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"context"
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social/client"
	"strconv"
)

const (
	Base = "https://api.dribbble.com/"
)

type Client struct {
	User  *UserService
	Shots *ShotService
}

// NewClient returns a new Dribbble Client.
func NewClient(ctx context.Context, c *oauth.Credentials, token *oauth2.Token) *Client {
	cl := client.NewHttpClient().Base(Base)
	auther := oauth2.NewOAuth(ctx, c, token, cl)
	return &Client{
		User:  newUserService(auther),
		Shots: newShotService(auther),
	}
}

func (d *Client) GoSocialUser() (*models.SocialUser, error) {
	s, err := d.Shots.DribbbleShots()
	if err != nil {
		return nil, err
	}

	u, err := d.User.UserCredentials()
	if err != nil {
		return nil, err
	}

	goSocial := models.SocialUser{
		Username:     u.Login,
		Name:         u.Name,
		UserId:       strconv.Itoa(u.ID),
		ContentCount: int64(len(*s)),
		Verified:     u.Pro,
		AvatarUrl:    u.AvatarURL,
		Followers:    u.FollowersCount,
		Url:          u.HTMLURL,
	}

	return &goSocial, nil
}
