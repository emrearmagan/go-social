/*
github.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package github

import (
	"go-social/models"
	"go-social/social/oauth/oauth2"
	"strconv"
	"strings"
)

const (
	Base                = "https://api.github.com/"
	AuthorizationPrefix = "token " // trailing space is required
)

type Client struct {
	User     *UserService
	Follower *FollowerService
}

// NewClient returns a new GitHub Client.
func NewClient(oauth *oauth2.OAuth2) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		User:     newUserService(oauth.New()),
		Follower: newFollowerService(oauth.New()),
	}
}

func (g *Client) GoSocialUser() (*models.SocialUser, error) {
	u, err := g.User.UserCredentials()
	if err != nil {
		return nil, err
	}

	pro := strings.ToLower(u.Plan.Name) == "pro"
	repos := u.PublicRepos + u.TotalPrivateRepos
	goSocial := models.SocialUser{
		Username:     u.Login,
		Name:         u.Name,
		UserId:       strconv.Itoa(u.ID),
		ContentCount: int64(repos),
		Verified:     pro,
		AvatarUrl:    u.AvatarURL,
		Followers:    u.Followers,
		Following:    &u.Following,
		Url:          u.HTMLURL,
	}

	return &goSocial, nil
}
