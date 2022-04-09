/*
dribbble.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package dribbble

import (
	"github.com/emrearmagan/go-social/models"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"strconv"
)

const (
	Base                = "https://api.dribbble.com/"
	AuthorizationPrefix = "Bearer " // trailing space is required
)

type Client struct {
	User  *UserService
	Shots *ShotService
}

// NewClient returns a new Dribbble Client.
func NewClient(oauth *oauth2.OAuth2) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		User:  newUserService(oauth),
		Shots: newShotService(oauth),
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
