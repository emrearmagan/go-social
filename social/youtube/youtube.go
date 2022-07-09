/*
youtube.go
Created at 09.07.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package youtube

import (
	"context"
	"github.com/emrearmagan/go-social/oauth"
	"github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
	"github.com/emrearmagan/go-social/social/client"
	"strings"
)

const (
	APIBase = "https://youtube.googleapis.com/"

	ClientHeaderName = "Client-Id"

	RefreshRevokeBase = "https://oauth2.googleapis.com/"
	RefreshPath       = "/token"
	RevokePath        = "/revoke"
)

type Client struct {
	oauth2 *oauth2.OAuth2

	User    *UserService
	Channel *ChannelService
	Search  *SearchService
}

// NewClient returns a new Youtube Client.
func NewClient(ctx context.Context, c *oauth.Credentials, token *oauth2.Token) *Client {
	// YouTube requires the client id to be in the header. At least for the endpoints implemented here
	cl := client.NewHttpClient().Base(APIBase)
	cl.Add(ClientHeaderName, c.ConsumerKey)
	auther := oauth2.NewOAuth(ctx, c, token, cl)
	return &Client{
		oauth2:  auther,
		User:    newUserService(auther),
		Channel: newChannelService(auther),
		Search:  newSearchService(auther),
	}
}

// RefreshToken a new access token can be generated by supplying the refresh token originally obtained during
// authorization code exchange.
// https://developers.google.com/youtube/v3/guides/auth/installed-apps#offline
func (c *Client) RefreshToken() (*oauth2.OAuthRefreshResponse, error) {
	oauthResp := new(OAuth2Response)
	apiError := new(APIError)

	// Youtube requires the client id to be in the url of the request.
	rclient := c.oauth2.Client()
	rclient.AddQuery(struct {
		ClientId string `url:"client_id"`
	}{
		ClientId: c.oauth2.Credentials().ConsumerKey,
	})
	oauth := c.oauth2.NewClient(rclient)
	err := oauth.RefreshToken(RefreshRevokeBase, RefreshPath, oauthResp, apiError)
	c.oauth2.UpdateToken(oauth2.NewToken(oauthResp.AccessToken, c.oauth2.Token().RefreshToken))
	return &oauth2.OAuthRefreshResponse{
		Token: oauth2.Token{
			Token:        oauthResp.AccessToken,
			RefreshToken: c.oauth2.Token().RefreshToken,
		},
		TokenType: oauthResp.TokenType,
		ExpiresIn: oauthResp.ExpiresIn,
		Scope:     strings.Fields(oauthResp.Scope),
	}, social.CheckError(err)
}

// Revoke If your app no longer needs an access token.
// If the revocation is successfully processed, then the HTTP status code of the response is 200.
// For error conditions, an HTTP status code 400 is returned along with an error code.
// https://developers.google.com/youtube/v3/guides/auth/installed-apps#tokenrevoke
func (c *Client) Revoke() error {
	apiError := new(APIError)

	// YouTube requires the token to be in the url of the request.
	// The token can be an access token or a refresh token. If the token is an access token and it has a corresponding refresh token, the refresh token will also be revoked.
	rclient := c.oauth2.Client()
	rclient.AddQuery(struct {
		Token string `url:"token"`
	}{
		Token: c.oauth2.Token().Token,
	})
	oauth := c.oauth2.NewClient(rclient)
	err := oauth.RevokeToken(RefreshRevokeBase, RevokePath, nil, apiError)
	return social.CheckError(err)
}

type OAuth2Response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	IdToken     string `json:"id_token"`
}
