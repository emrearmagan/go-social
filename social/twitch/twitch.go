/*
twitch.go
Created at 19.06.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package twitch

import (
	"context"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"github.com/emrearmagan/go-social/social/client"
)

type Client struct {
	oauth2 *oauth2.OAuth2

	User       *UserService
	Subscriber *SubscriberService
	Follower   *FollowerService
}

const (
	APIBase = "https://api.twitch.tv/"

	ClientHeaderName = "Client-Id"

	RefreshRevokeBase = "https://id.twitch.tv/"
	RefreshPath       = "/oauth2/token"
	RevokePath        = "/oauth2/revoke"
)

// NewClient returns a new Twitter Client.
func NewClient(ctx context.Context, c *oauth.Credentials, token *oauth2.Token) *Client {
	// Twitch requires the client id to be in the header. At least for the endpoints implemented here
	cl := client.NewHttpClient().Base(APIBase)
	cl.Add(ClientHeaderName, c.ConsumerKey)
	auther := oauth2.NewOAuth(ctx, c, token, cl)
	return &Client{
		oauth2:     auther,
		User:       newUserService(auther),
		Subscriber: newSubscriberService(auther),
		Follower:   newFollowerService(auther),
	}
}

// RefreshToken a new access token can be generated by supplying the refresh token originally obtained during
// authorization code exchange.
// https://dev.twitch.tv/docs/authentication/refresh-tokens
func (c *Client) RefreshToken() (*oauth2.OAuthRefreshResponse, error) {
	oauthResp := new(OAuth2Response)
	apiError := new(APIError)

	// Twitch requires the client id and secret to be in the body of the request.
	rclient := c.oauth2.Client()
	rclient.AddQuery(struct {
		ClientId     string `url:"client_id"`
		ClientSecret string `url:"client_secret"`
	}{
		ClientId:     c.oauth2.Credentials().ConsumerKey,
		ClientSecret: c.oauth2.Credentials().ConsumerSecret,
	})
	oauth := c.oauth2.NewClient(rclient)

	err := oauth.RefreshToken(RefreshRevokeBase, RefreshPath, oauthResp, apiError)
	c.oauth2.UpdateToken(oauth2.NewToken(oauthResp.AccessToken, oauthResp.RefreshToken))
	return &oauth2.OAuthRefreshResponse{
		Token: oauth2.Token{
			Token:        oauthResp.AccessToken,
			RefreshToken: oauthResp.RefreshToken,
		},
		TokenType: oauthResp.TokenType,
		ExpiresIn: oauthResp.ExpiresIn,
		Scope:     oauthResp.Scope,
	}, social.CheckError(err)
}

// Revoke If your app no longer needs an access token.
// Returns 400 Bad Request if the client ID is valid but the access token is not
// Returns 404 Not Found if the client ID is not valid.
func (c *Client) Revoke() error {
	apiError := new(APIError)

	// Twitch requires the client id and secret to be in the body of the request.
	rclient := c.oauth2.Client()
	rclient.AddQuery(struct {
		ClientId string `url:"client_id"`
		Token    string `url:"token"`
	}{
		ClientId: c.oauth2.Credentials().ConsumerKey,
		Token:    c.oauth2.Token().Token,
	})

	oauth := c.oauth2.NewClient(rclient)
	err := oauth.RevokeToken(RefreshRevokeBase, RevokePath, nil, apiError)
	return social.CheckError(err)
}

type OAuth2Response struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}
