/*
token.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth2

// Token represents an OAuth1 AccessToken (Token credentials) and secret
type Token struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewToken returns a new OAuth1 Token
func NewToken(token, refreshToken string) *Token {
	return &Token{
		Token:        token,
		RefreshToken: refreshToken,
	}
}
