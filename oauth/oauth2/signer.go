/*
signer.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth2

import (
	"fmt"
	"strings"
)

// A Signer signs request to create a signed 0Auth2 request
type Signer interface {
	Name() string
	// AuthSigningParams returns the OAuth2 header parameters for the given credentials
	// See https://tools.ietf.org/html/rfc6749
	AuthSigningParams() map[string]string
	OAuthParams(token string) map[string]string
}

const (
	BasicAuthorizationPrefix  = "Basic "  // trailing space is required
	BearerAuthorizationPrefix = "Bearer " // trailing space is required
	AuthorizationHeaderName   = "Authorization"
	ContentLengthHeaderName   = "Content-Length"
	ContentTypeHeaderName     = "Content-Type"
)

// A BasicSigner signs request with an basic header prefix
type BasicSigner struct {
	ConsumerKey    string
	ConsumerSecret string
}

func (b BasicSigner) Name() string {
	return BasicAuthorizationPrefix
}

func (b BasicSigner) AuthSigningParams() map[string]string {
	return map[string]string{
		AuthorizationHeaderName: b.authorizationHeaderValue(),
		ContentTypeHeaderName:   "application/x-www-form-urlencoded",
	}
}

func (b BasicSigner) OAuthParams(token string) map[string]string {
	header := []string{BasicAuthorizationPrefix, token}
	return map[string]string{
		AuthorizationHeaderName: strings.Join(header, ""),
		ContentTypeHeaderName:   "application/json",
	}
}

func (b BasicSigner) authorizationHeaderValue() string {
	return BasicAuthorizationPrefix + Base64Enc(fmt.Sprintf("%s:%s", b.ConsumerKey, b.ConsumerSecret))
}

// A BearerSigner signs request with a bearer header prefix
type BearerSigner struct {
	ConsumerKey    string
	ConsumerSecret string
}

func (b BearerSigner) Name() string {
	return BearerAuthorizationPrefix
}

func (b BearerSigner) OAuthParams(token string) map[string]string {
	header := []string{BearerAuthorizationPrefix, token}

	return map[string]string{
		AuthorizationHeaderName: strings.Join(header, ""),
		ContentTypeHeaderName:   "application/json",
	}
}

// AuthSigningParams returns the OAuth2 header parameters for the given credentials
// See https://tools.ietf.org/html/rfc6749
func (b BearerSigner) AuthSigningParams() map[string]string {
	return map[string]string{
		AuthorizationHeaderName: b.authorizationHeaderValue(),
		ContentTypeHeaderName:   "application/x-www-form-urlencoded",
	}
}
func (b BearerSigner) authorizationHeaderValue() string {
	return BearerAuthorizationPrefix + Base64Enc(fmt.Sprintf("%s:%s", b.ConsumerKey, b.ConsumerSecret))
}
