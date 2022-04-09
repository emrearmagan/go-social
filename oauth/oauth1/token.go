/*
token.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth1

// Token represents an OAuth1 AccessToken (token credentials) and secret
type Token struct {
	Token       string `json:"access_token"`
	TokenSecret string `json:"token_secret"`
}

// NewToken returns a new OAuth1 Token
func NewToken(token, tokenSecret string) *Token {
	return &Token{
		Token:       token,
		TokenSecret: tokenSecret,
	}
}
