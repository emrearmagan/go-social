/*
oauth2.go
Created at 08.04.22 by emrearmagan
Copyright © go-social. All rights reserved.
*/

package oauth2

import (
	"context"
	"errors"
	"fmt"
	"go-social/social"
	"go-social/social/oauth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	AuthorizationHeaderName = "Authorization"
	ContentTypeHeaderName   = "Content-Type"
	ContentLengthHeaderName = "Content-Length"

	authorizationPrefix = "Basic " // trailing space is required
)

type OAuth2 struct {
	ctx         context.Context
	credentials *oauth.Credentials
	token       *Token
	client      *social.HttpClient
}

// oauthParams returns the OAuth2 header parameters for the given credentials
// See https://tools.ietf.org/html/rfc6749
func (a *OAuth2) oAuthParams() map[string]string {
	return map[string]string{
		AuthorizationHeaderName: authorizationHeaderValue(a.credentials),
		ContentTypeHeaderName:   "application/x-www-form-urlencoded",
	}
}

func NewOAuth2(ctx context.Context, c *oauth.Credentials, token *Token) *OAuth2 {
	return &OAuth2{
		ctx:         ctx,
		credentials: c,
		token:       token,
		client:      social.NewClient(),
	}
}

func (a *OAuth2) New() *OAuth2 {
	return &OAuth2{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      a.client.New(),
	}
}

func (a *OAuth2) NewClient(client *social.HttpClient) *OAuth2 {
	return &OAuth2{
		ctx:         a.ctx,
		credentials: a.credentials,
		token:       a.token,
		client:      client,
	}
}

func (a *OAuth2) SignRequest(req *http.Request) (*http.Request, error) {
	if a.credentials.ConsumerKey == "" {
		return nil, errors.New("OAuth2: provide valid credentials")
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", a.token.RefreshToken)

	req = req.WithContext(a.ctx)
	req.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))

	for k, v := range a.oAuthParams() {
		req.Header.Set(k, v)
	}

	req.Header.Add(ContentLengthHeaderName, strconv.Itoa(len(data.Encode())))

	return req, nil
}

func (a *OAuth2) AddHeader(headerPrefix string) (*http.Request, error) {
	req, err := a.client.Request()
	if err != nil {
		return nil, err
	}
	req = req.WithContext(a.ctx)
	header := []string{headerPrefix, a.token.Token}
	req.Header.Set(AuthorizationHeaderName, strings.Join(header, ""))

	return req, nil
}

func (a *OAuth2) Client() *social.HttpClient {
	return a.client
}

func authorizationHeaderValue(c *oauth.Credentials) string {
	return authorizationPrefix + base64Enc(fmt.Sprintf("%s:%s", c.ConsumerKey, c.ConsumerSecret))
}
