/*
spotify.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import "go-social/social/oauth/oauth2"

const (
	Base                = "https://api.spotify.com/"
	AuthorizationPrefix = "Bearer " // trailing space is required
)

type Client struct {
	Account *AccountService
	User    *UserService
}

// NewClient returns a new Dribbble Client.
func NewClient(oauth *oauth2.OAuth2) *Client {
	oauth = oauth.NewClient(oauth.Client().Base(Base))
	oauth.AuthorizationPrefix = AuthorizationPrefix
	return &Client{
		Account: newAccountService(oauth),
		User:    newUserService(oauth),
	}
}
