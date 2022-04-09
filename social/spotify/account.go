/*
account.go
Created at 09.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package spotify

import (
	oauth22 "github.com/emrearmagan/go-social/oauth/oauth2"
	"github.com/emrearmagan/go-social/social"
)

const (
	RefreshBase = "https://accounts.spotify.com"
	RefreshPath = "/api/token/"
)

// AccountService provides methods for account information
type AccountService struct {
	oauth2 *oauth22.OAuth2
}

// newAccountService returns a new Spotify UserService.
func newAccountService(oauth2 *oauth22.OAuth2) *AccountService {
	return &AccountService{
		oauth2: oauth2,
	}
}

// RefreshToken a new access token can be generated by supplying the refresh token originally obtained during
// authorization code exchange.
// https://developer.spotify.com/documentation/general/guides/authorization-guide/
func (a *AccountService) RefreshToken() (*oauth22.OAuthRefreshResponse, error) {
	oauthResp := new(OAuth2Response)
	apiError := new(APIError)

	err := a.oauth2.RefreshToken(RefreshBase, RefreshPath, oauthResp, apiError)
	a.oauth2.UpdateToken(oauth22.NewToken(oauthResp.AccessToken, a.oauth2.Token().RefreshToken))
	return &oauth22.OAuthRefreshResponse{
		Token: oauth22.Token{
			Token:        oauthResp.AccessToken,
			RefreshToken: a.oauth2.Token().RefreshToken,
		},
		TokenType: oauthResp.TokenType,
		ExpiresIn: oauthResp.ExpiresIn,
		Scope:     oauthResp.Scope,
	}, social.CheckError(err)
}

type OAuth2Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
