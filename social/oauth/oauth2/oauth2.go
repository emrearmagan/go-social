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

	BasicAuthorizationPrefix = "Basic " // trailing space is required
)

type OAuth2 struct {
	ctx         context.Context
	credentials *oauth.Credentials
	token       *Token
	client      *social.HttpClient

	AuthorizationPrefix string
}

func NewOAuth2(ctx context.Context, c *oauth.Credentials, token *Token) *OAuth2 {
	return &OAuth2{
		ctx:                 ctx,
		credentials:         c,
		token:               token,
		client:              social.NewClient(),
		AuthorizationPrefix: BasicAuthorizationPrefix,
	}
}

func (a *OAuth2) New() *OAuth2 {
	return &OAuth2{
		ctx:                 a.ctx,
		credentials:         a.credentials,
		token:               a.token,
		client:              a.client.New(),
		AuthorizationPrefix: a.AuthorizationPrefix,
	}
}

func (a *OAuth2) NewClient(client *social.HttpClient) *OAuth2 {
	return &OAuth2{
		ctx:                 a.ctx,
		credentials:         a.credentials,
		token:               a.token,
		client:              client,
		AuthorizationPrefix: a.AuthorizationPrefix,
	}
}

func (a *OAuth2) Get(path string, resp interface{}, apiError social.Errors, params interface{}) error {
	client := a.client.AddQuery(params).Get(path)

	req, err := a.client.Request()
	if err != nil {
		return err
	}
	req = req.WithContext(a.ctx)
	for k, v := range a.oAuthParams() {
		req.Header.Set(k, v)
	}

	httpResp, err := client.Do(req, resp, apiError.ErrorDetail())
	if err != nil {
		return err
	}

	if httpResp != nil {
		if code := httpResp.StatusCode; code >= 300 {
			apiError.SetStatus(httpResp.StatusCode)
		}
	}

	return social.RelevantError(err, apiError)
}

func (a *OAuth2) oAuthParams() map[string]string {
	header := []string{a.AuthorizationPrefix, a.token.Token}

	return map[string]string{
		AuthorizationHeaderName: strings.Join(header, ""),
		ContentTypeHeaderName:   "application/json",
	}
}

func (a *OAuth2) AddCustomHeaders(key, value string) {
	a.client = a.client.Set(key, value)
}

//----------------
func (a *OAuth2) signRequest(req *http.Request) (*http.Request, error) {
	if a.credentials.ConsumerKey == "" {
		return nil, errors.New("OAuth2: provide valid credentials")
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", a.token.RefreshToken)

	req = req.WithContext(a.ctx)
	req.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))

	for k, v := range a.oAuthSigningParams() {
		req.Header.Set(k, v)
	}

	req.Header.Add(ContentLengthHeaderName, strconv.Itoa(len(data.Encode())))

	return req, nil
}

func (a *OAuth2) RefreshToken(refreshBase string, path string) (*http.Request, error) {
	client := a.Client().Base(refreshBase).Post(path)

	req, err := client.Request()
	if err != nil {
		return nil, err
	}

	req = req.WithContext(a.ctx)
	req, err = a.signRequest(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (a *OAuth2) Client() *social.HttpClient {
	return a.client
}

// oauthParams returns the OAuth2 header parameters for the given credentials
// See https://tools.ietf.org/html/rfc6749
func (a *OAuth2) oAuthSigningParams() map[string]string {
	return map[string]string{
		AuthorizationHeaderName: authorizationHeaderValue(a.credentials),
		ContentTypeHeaderName:   "application/x-www-form-urlencoded",
	}
}

func authorizationHeaderValue(c *oauth.Credentials) string {
	//TODO: Use the given autorization prefix
	return BasicAuthorizationPrefix + base64Enc(fmt.Sprintf("%s:%s", c.ConsumerKey, c.ConsumerSecret))
}
